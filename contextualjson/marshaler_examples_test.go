package contextualjson_test

import (
	"encoding/json"
	"fmt"

	"github.com/go-extras/go-kit/contextualjson"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email" marshalcontext:"private"`
	FirstName string `json:"first_name" marshalhandler:"FullName"`
	LastName  string `json:"last_name" marshalhandler:"FullName"`
}

func (u User) FullName(_ string) (value, jsonTag string) {
	return u.FirstName + " " + u.LastName, "full_name"
}

func ExampleMarshaler() {
	user := User{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	// Default serialization without context
	marshaler := contextualjson.NewMarshaler(user, "")
	data, err := marshaler.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var defaultOutput map[string]any
	err = json.Unmarshal(data, &defaultOutput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %v, username: %s, full_name: %s\n", int(defaultOutput["id"].(float64)), defaultOutput["username"], defaultOutput["full_name"])

	// Serialization with the "private" context
	privateMarshaler := contextualjson.NewMarshaler(user, "private")
	privateData, err := privateMarshaler.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var privateOutput map[string]any
	err = json.Unmarshal(privateData, &privateOutput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("email: %s, id: %v, username: %s, full_name: %s\n", privateOutput["email"], int(privateOutput["id"].(float64)), privateOutput["username"], privateOutput["full_name"])

	// Output:
	// id: 1, username: johndoe, full_name: John Doe
	// email: johndoe@example.com, id: 1, username: johndoe, full_name: John Doe
}
