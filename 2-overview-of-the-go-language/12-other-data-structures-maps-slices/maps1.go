package main

import "fmt"

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	fmt.Println(myMap)

	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Fido"

	fmt.Println(myMap)

	myOtherMap := make(map[string]int)
	myOtherMap["First"] = 1
	myOtherMap["Second"] = 2

	fmt.Println(myOtherMap)
}
