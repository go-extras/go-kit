// Package contextualjson provides a JSON marshaler that allows specifying a context for the serialization of struct fields.
//
// To use, create a Marshaler instance with the data to be serialized, a context string, and any options (e.g., marshalcontext tag, marshalhandler tag).
// Then call MarshalJSON on the Marshaler instance to serialize the data with the given context.
//
// Example usage:
//
//	type Person struct {
//	    Name     string `json:"name"`
//	    Age      int    `json:"age"`
//	    Password string `json:"password" serialize:"admin"`
//	}
//
//	// serialize only the Name and Age fields
//	m := NewMarshaler(Person{Name: "Alice", Age: 30, Password: "password"}, "user")
//	bytes, err := json.Marshal(m)
//
//	// serialize all fields, including Password (only when context is "admin")
//	m = NewMarshaler(Person{Name: "Alice", Age: 30, Password: "password"}, "admin")
//	bytes, err = json.Marshal(m)
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package contextualjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	defaultMarshalContextTag = "marshalcontext"
	defaultMarshalHandlerTag = "marshalhandler"
)

// MarshalerOption is a functional option type for configuring a Marshaler instance.
type MarshalerOption func(*Marshaler)

// Marshaler is a JSON marshaler that allows specifying a context for the serialization of struct fields.
type Marshaler struct {
	data              any
	context           string
	marshalContextTag string
	marshalHandlerTag string
}

// WithMarshalContextTag returns a MarshalerOption that sets the tag to use for context-aware field serialization.
// The default is "marshalcontext".
func WithMarshalContextTag(tag string) MarshalerOption {
	return func(m *Marshaler) {
		m.marshalContextTag = tag
	}
}

// WithMarshalHandlerTag returns a MarshalerOption that sets the tag to use for marshaling a field with a custom function.
// The default is "marshalhandler".
func WithMarshalHandlerTag(tag string) MarshalerOption {
	return func(m *Marshaler) {
		m.marshalHandlerTag = tag
	}
}

// NewMarshaler returns a new Marshaler instance with the given data, context, and options.
func NewMarshaler(data any, context string, opts ...MarshalerOption) *Marshaler {
	m := &Marshaler{
		data:              data,
		context:           context,
		marshalContextTag: defaultMarshalContextTag,
		marshalHandlerTag: defaultMarshalHandlerTag,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// MarshalJSON marshals the data with the context specified in the Marshaler instance.
func (m *Marshaler) MarshalJSON() ([]byte, error) {
	return MarshalJSONWithContext(m.data, m.context, m.marshalContextTag, m.marshalHandlerTag)
}

// MarshalJSONWithContext marshals the data with the given context.
func MarshalJSONWithContext(data any, context, contextTag, handlerTag string) ([]byte, error) {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	contextValue := reflect.ValueOf(context)

	var err error

	switch value.Kind() {
	case reflect.Struct:
		fields := make(map[string]any)
		for i := 0; i < value.NumField(); i++ {
			err = processFields(value, contextValue, i, context, contextTag, handlerTag, fields)
			if err != nil {
				return nil, err
			}
		}
		return json.Marshal(fields)
	case reflect.Slice, reflect.Array:
		elems := make([]any, value.Len())
		for i := 0; i < value.Len(); i++ {
			elem := value.Index(i).Interface()
			elems[i], err = MarshalJSONWithContext(elem, context, contextTag, handlerTag)
			if err != nil {
				return nil, err
			}
		}
		return json.Marshal(elems)
	case reflect.Map:
		keys := value.MapKeys()
		mp := make(map[string]any)
		for _, k := range keys {
			keyStr := fmt.Sprintf("%v", k.Interface())
			val := value.MapIndex(k).Interface()
			mp[keyStr], err = MarshalJSONWithContext(val, context, contextTag, handlerTag)
			if err != nil {
				return nil, err
			}
		}
		return json.Marshal(mp)
	default:
		return json.Marshal(data)
	}
}

func processFields(value, contextValue reflect.Value, i int, context, contextTag, handlerTag string, fields map[string]any) error {
	field := value.Type().Field(i)
	if field.Anonymous {
		// iterate over the embedded struct's fields
		for j := 0; j < value.Field(i).NumField(); j++ {
			err := processFields(value.Field(i), contextValue, j, context, contextTag, handlerTag, fields)
			if err != nil {
				return err
			}
		}
		return nil
	}

	fieldValue := value.Field(i).Interface()
	marshalCtx := field.Tag.Get(contextTag)
	if marshalCtx != "" && marshalCtx != context {
		return nil
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		jsonTag = strings.ToLower(field.Name)
	}

	err := maybeGetJSONTagFromHandler(field, handlerTag, value, contextValue, &fieldValue, &jsonTag)
	if err != nil {
		return err
	}

	if jsonTag != "-" {
		fields[jsonTag] = fieldValue
	}

	return nil
}

func maybeGetJSONTagFromHandler(field reflect.StructField, handlerTag string, value, contextValue reflect.Value, fieldValue *any, jsonTag *string) error {
	if handlerVal := field.Tag.Get(handlerTag); handlerVal != "" {
		handlerName, newFieldName := parseCtxHandlerTag(handlerVal)
		err := getJSONTagFromHandler(value, contextValue, newFieldName, handlerName, fieldValue, jsonTag)
		if err != nil {
			return err
		}
	}
	return nil
}

func getJSONTagFromHandler(value, contextValue reflect.Value, newFieldName, handlerName string, fieldValue *any, jsonTag *string) error {
	// Value can be a pointer to a struct or a struct.
	// Value can have a method with a pointer receiver or a value receiver.
	// Attempt to find a method with a pointer receiver first.
	// Then attempt to find a method with a value receiver.
	// If both fail, return an error.

	handlerFunc := findHandlerFunc(value, handlerName)
	if handlerFunc.IsValid() { // first, check if we have the function
		if !isHandlerFuncValid(handlerFunc) { // then check if it's valid
			return fmt.Errorf("invalid handler func signature, must be func(marshalCtx string) (value any, jsonTag string) OR func(marshalCtx string) (value any)")
		}

		handlerResult := handlerFunc.Call([]reflect.Value{contextValue})
		*fieldValue = handlerResult[0].Interface()
		if newFieldName != "" {
			*jsonTag = newFieldName
		}
		if len(handlerResult) <= 1 {
			return nil
		}

		var ok bool
		// take jsonTag from the second return value
		if !handlerResult[1].CanInterface() {
			return fmt.Errorf("invalid json tag returned by handler %+v", handlerResult[1])
		}
		*jsonTag, ok = handlerResult[1].Interface().(string)
		if !ok {
			return fmt.Errorf("invalid json tag returned by handler %+v", handlerResult[1].Interface())
		}

		return nil
	}

	if newFieldName != "" {
		*jsonTag = newFieldName
	}

	return nil
}

func parseCtxHandlerTag(tag string) (handlerName, newFieldName string) {
	parts := strings.Split(tag, ",")
	handlerName = parts[0]
	if len(parts) > 1 {
		newFieldName = parts[1]
	}
	return handlerName, newFieldName
}

func findHandlerFunc(value reflect.Value, handlerName string) reflect.Value {
	handlerFunc := value.MethodByName(handlerName)
	if handlerFunc.IsValid() {
		return handlerFunc
	}

	if value.CanAddr() {
		handlerFunc = value.Addr().MethodByName(handlerName)
		if handlerFunc.IsValid() {
			return handlerFunc
		}
	}

	if value.Kind() == reflect.Ptr {
		handlerFunc = value.Elem().MethodByName(handlerName)
		if handlerFunc.IsValid() {
			return handlerFunc
		}
	}

	pointer := reflect.New(value.Type())
	pointer.Elem().Set(value)
	handlerFunc = pointer.MethodByName(handlerName)
	if handlerFunc.IsValid() {
		return handlerFunc
	}

	return reflect.Value{}
}

func isHandlerFuncValid(handlerFunc reflect.Value) bool {
	// Check if the method signature has a single 'string' argument
	if handlerFunc.Type().NumIn() != 1 || handlerFunc.Type().In(0).Kind() != reflect.String {
		return false
	}

	// Check if the method signature matches func(string) (any, string)
	if handlerFunc.Type().NumOut() == 2 && handlerFunc.Type().Out(1).Kind() == reflect.String {
		return true
	}

	// Check if the method signature matches func(string) any
	if handlerFunc.Type().NumOut() == 1 {
		return true
	}

	return false
}
