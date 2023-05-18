package main

import (
	"fmt"
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)

	otherVar := helpers.SomeType{
		TypeName:   "John",
		TypeNumber: 25,
	}

	fmt.Printf("%#v\n", otherVar)
}
