package main

import "fmt"

func main() {
	year := 1988
	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Println("Leap year")
		} else {
			fmt.Println("Not a leap year.")
		}

	} else {
		fmt.Println("Not a leap year.")
	}
}
