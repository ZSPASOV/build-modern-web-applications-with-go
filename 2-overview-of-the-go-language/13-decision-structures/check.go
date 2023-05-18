package main

import "fmt"

func main() {
	animal := "asd"

	switch animal {
	case "dog":
		fmt.Println("dog")
	case "cat":
		fmt.Println("cat")
	default:
		fmt.Println("other")
	}
}
