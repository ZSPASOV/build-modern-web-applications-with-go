package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	for j := 0; j < 7; j++ {
		log.Println(j)
	}

	animals := []string{"dog", "cat", "fish", "cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

	people := map[string]int{
		"Georgi": 30,
		"Ivan":   34,
		"Niki":   45,
		"Todor":  50,
	}

	for name, age := range people {
		log.Println(name)
		log.Println(age)
	}

	sentence := "Today is a good day"

	for _, letter := range sentence {
		log.Println(string(letter))
	}
}
