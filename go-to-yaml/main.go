package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Food struct {
	Price int
	Name  string
	Cool  bool
	Types []string
}

func main() {
	f := Food {
		Price: 10,
		Name: "Cuscuz",
		Cool: true,
		Types: []string{"recheado", "molho"},
	}

	out, err :=yaml.Marshal(f)
	if err != nil {
			log.Fatal(err)

	}

	fmt.Println(string(out))
}