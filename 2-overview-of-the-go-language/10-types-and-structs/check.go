package main

import (
	"fmt"
	"log"
)

var s = "seven"

func main() {
	var s2 = "six"

	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	a, b := saySomething("xxx")
	fmt.Println(a)
	fmt.Println(b)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
