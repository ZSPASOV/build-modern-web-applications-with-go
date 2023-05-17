package main

import "fmt"

type person struct {
	age    int
	weight int
}

func (p *person) calc() int {
	return p.age * p.weight
}

func main() {
	joro := person{
		age:    30,
		weight: 100,
	}
	fmt.Println(joro.calc())
}
