package contextualjson_test

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/go-extras/go-kit/contextualjson"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address" marshalcontext:"private"`
	Custom  string `json:"-" marshalhandler:"CustomHandler"`
}

func (p *Person) CustomHandler(_ string) (value string, jsonTag string) {
	// custom handling logic goes here
	return "custom value", "custom_name"
}

type PersonNoHandler Person

type PersonCustomTags struct {
	Person
	Address string `json:"address" customctx:"private"`
	Custom  string `json:"-" customfn:"CustomHandler"`
}

func (p *PersonCustomTags) CustomHandler(_ string) (value string, jsonTag string) {
	// custom handling logic goes here
	return "custom value", "custom_name"
}

func TestMarshalJSONWithContext(t *testing.T) {
	person := PersonNoHandler{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
	}

	// Serialize the person object with the context "public".
	publicData, err := contextualjson.MarshalJSONWithContext(person, "public", "marshalcontext", "marshalhandler")
	if err != nil {
		t.Errorf("unexpected error while serializing data: %v", err)
	}

	// Verify that the publicData only contains the "name" and "age" fields.
	var expectedPublicData = map[string]any{
		"name": person.Name,
		"age":  person.Age,
	}
	expectedPublicDataJSON, _ := json.Marshal(expectedPublicData)
	expectedPublicDataLines := strings.Split(string(expectedPublicDataJSON), "\n")
	sort.Strings(expectedPublicDataLines)
	expectedPublicDataJSON = []byte(strings.Join(expectedPublicDataLines, "\n"))
	publicDataLines := strings.Split(string(publicData), "\n")
	sort.Strings(publicDataLines)
	publicData = []byte(strings.Join(publicDataLines, "\n"))
	if !reflect.DeepEqual(publicData, expectedPublicDataJSON) {
		t.Errorf("unexpected publicData:\nGot: %s\nExpected: %s", string(publicData), string(expectedPublicDataJSON))
	}

	// Serialize the person object with the context "private".
	privateData, err := contextualjson.MarshalJSONWithContext(person, "private", "marshalcontext", "marshalhandler")
	if err != nil {
		t.Errorf("unexpected error while serializing data: %v", err)
	}

	// Verify that the privateData contains all the fields.
	var expectedPrivateData = map[string]any{
		"name":    person.Name,
		"age":     person.Age,
		"address": person.Address,
	}
	expectedPrivateDataJSON, _ := json.Marshal(expectedPrivateData)
	expectedPrivateDataLines := strings.Split(string(expectedPrivateDataJSON), "\n")
	sort.Strings(expectedPrivateDataLines)
	expectedPrivateDataJSON = []byte(strings.Join(expectedPrivateDataLines, "\n"))
	privateDataLines := strings.Split(string(privateData), "\n")
	sort.Strings(privateDataLines)
	privateData = []byte(strings.Join(privateDataLines, "\n"))
	if !reflect.DeepEqual(privateData, expectedPrivateDataJSON) {
		t.Errorf("unexpected privateData:\nGot: %s\nExpected: %s", string(privateData), string(expectedPrivateDataJSON))
	}
}

func TestMarshaler_MarshalJSON(t *testing.T) {
	person := PersonNoHandler{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
	}

	// Create a new Marshaler with the person object and the context "public".
	marshaler := contextualjson.NewMarshaler(person, "public")

	// Serialize the marshaler object and verify the result.
	data, err := json.Marshal(marshaler)
	if err != nil {
		t.Errorf("unexpected error while serializing data: %v", err)
	}

	// Verify that the data only contains the "name" and "age" fields.
	var expectedData = map[string]any{
		"name": person.Name,
		"age":  person.Age,
	}
	expectedDataJSON, _ := json.Marshal(expectedData)
	expectedDataLines := strings.Split(string(expectedDataJSON), "\n")
	sort.Strings(expectedDataLines)
	expectedDataJSON = []byte(strings.Join(expectedDataLines, "\n"))
	dataLines := strings.Split(string(data), "\n")
	sort.Strings(dataLines)
	data = []byte(strings.Join(dataLines, "\n"))
	if !reflect.DeepEqual(data, expectedDataJSON) {
		t.Errorf("unexpected data:\nGot: %s\nExpected: %s", string(data), string(expectedDataJSON))
	}
}

func TestMarshaler_MarshalJSON_CustomTags(t *testing.T) {
	person := PersonCustomTags{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Address: "123 Main St",
	}

	// Create a new Marshaler with the person object and the context "public".
	marshaler := contextualjson.NewMarshaler(person, "public")

	// Serialize the marshaler object and verify the result.
	data, err := json.Marshal(marshaler)
	if err != nil {
		t.Errorf("unexpected error while serializing data: %v", err)
	}

	// Verify that the data only contains the "name" and "age" fields.
	var expectedData = map[string]any{
		"name": person.Name,
		"age":  person.Age,
	}
	expectedDataJSON, _ := json.Marshal(expectedData)
	expectedDataLines := strings.Split(string(expectedDataJSON), "\n")
	sort.Strings(expectedDataLines)
	expectedDataJSON = []byte(strings.Join(expectedDataLines, "\n"))
	dataLines := strings.Split(string(data), "\n")
	sort.Strings(dataLines)
	data = []byte(strings.Join(dataLines, "\n"))
	if !reflect.DeepEqual(data, expectedDataJSON) {
		t.Errorf("unexpected data:\nGot: %s\nExpected: %s", string(data), string(expectedDataJSON))
	}
}

func TestMarshalJSONWithContextCustomValue(t *testing.T) {
	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
		Custom:  "custom value",
	}

	// Create a new Marshaler with the person object and the context "public".
	marshaler := contextualjson.NewMarshaler(person, "public")

	// Serialize the person object with the context "public".
	publicData, err := json.Marshal(marshaler)
	if err != nil {
		t.Errorf("unexpected error while serializing data: %v", err)
	}

	// Verify that the publicData only contains the "name", "age", and "custom_name" fields.
	var expectedPublicData = map[string]any{
		"name":        person.Name,
		"age":         person.Age,
		"custom_name": "custom value",
	}
	expectedPublicDataJSON, _ := json.Marshal(expectedPublicData)
	expectedPublicDataLines := strings.Split(string(expectedPublicDataJSON), "\n")
	sort.Strings(expectedPublicDataLines)
	expectedPublicDataJSON = []byte(strings.Join(expectedPublicDataLines, "\n"))
	publicDataLines := strings.Split(string(publicData), "\n")
	sort.Strings(publicDataLines)
	publicData = []byte(strings.Join(publicDataLines, "\n"))
	if !reflect.DeepEqual(publicData, expectedPublicDataJSON) {
		t.Errorf("unexpected publicData:\nGot: %s\nExpected: %s", string(publicData), string(expectedPublicDataJSON))
	}
}
