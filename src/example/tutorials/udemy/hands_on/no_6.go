package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{
		"Alejo",
		"Covian",
		[]string{"carrots", "oranges", "peaches", "strawberries"},
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}
}
