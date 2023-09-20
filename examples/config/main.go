package main

import (
	"fmt"
	"github.com/VerTox/graphql-codegen-go/examples/config/internal"
)

func main() {
	person := internal.Person{
		FirstName: "",
		Lastname:  "",
		Age:       nil,
		Gender:    nil,
		Address:   nil,
	}
	fmt.Println(person)
}
