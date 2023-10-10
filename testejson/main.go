package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := `{"name": "John Doe", "age": 30, "email": "johndoe@example.com", "address": "123 Main St"}`

	var person Person
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded person:", person)
}
