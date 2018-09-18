package main

import "fmt"

func birthday(x int, y int) int {
	if x == 8 || y == 8 {
		return ("It's your birthday!")
	}
}

func main() {
	fmt.Println(birthday(8, 6))
}
