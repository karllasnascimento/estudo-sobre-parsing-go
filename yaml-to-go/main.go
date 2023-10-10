package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Food struct {
	Price int      `yaml:"Price"`
	Name  string   `yaml:"Name"`
	Cool  bool     `yaml:"Cool"`
	Types []string `yaml:"Types"`
}

func main() {
	//carreg arquivo e retorna info rm byte
	f, err := os.ReadFile("cuscuz.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// cria o Food para poder ser desempactptado
	var fo Food

	if err := yaml.Unmarshal(f, &fo); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n É bom demais!\n", fo)
}
