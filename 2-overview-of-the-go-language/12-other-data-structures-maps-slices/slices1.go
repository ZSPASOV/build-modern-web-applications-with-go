package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)

	var otherSlice []int
	otherSlice = append(otherSlice, 2)
	otherSlice = append(otherSlice, 1)
	otherSlice = append(otherSlice, 3)

	log.Println(otherSlice)

	sort.Ints(otherSlice)

	log.Println(otherSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[3:7])

	names := []string{"Kiro", "Yasen"}
	log.Println(names)
}
