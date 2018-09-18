package main

import "fmt"

//multiples of three, skipping -3 and -6
func main() {
	for i := -300; i <= 0; i++ {
		if i%3 == 0 {
			if i == -3 || i == -6 {
				continue
			} else {
				fmt.Println(i)
			}
		}
	}
}
