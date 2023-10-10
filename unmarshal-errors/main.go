package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	data, err := os.ReadFile("data.yaml")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	var person Person
	err = yaml.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("Erro ao deserializar o arquivo YAML:", err)
		return
	}
	fmt.Println("Nome:", person.Name)
	fmt.Println("Idade:", person.Age)
}
