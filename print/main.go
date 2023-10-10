// função que apenas realiza o print da informação

package main

import "fmt"

type Food struct {
	Price int
	Name  string
	Cool  bool
	Types []string
}

func main() {
	f := Food{
		Price: 20,
		Name:  "Cuscuz",
		Cool:  true,
		Types: []string{"recheado", "molho"},
	}

	fmt.Println(f)
}
