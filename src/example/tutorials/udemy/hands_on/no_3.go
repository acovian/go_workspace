package main

import "fmt"

func main() {
	s := []int{1, 2, 7, 8, 9}
	fmt.Println(s)

	for i := range s {
		fmt.Println(i)
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
