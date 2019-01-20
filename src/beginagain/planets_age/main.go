package main

import "fmt"

type Planet string

func main() {
	planetaryYear := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	var age int
	fmt.Println("How old are you?> ")
	fmt.Scanf("%d", &age)
	for planet, year := range planetaryYear {
		fmt.Println(planet, " - ", float64(age)/year)
	}

}
