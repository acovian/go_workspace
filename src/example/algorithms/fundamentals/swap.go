package main

import "fmt"

func main() {
	var myNumber = "42"
	var myName = "Alejo"
	var temp = myNumber
	myNumber = myName
	myName = temp
	fmt.Println(myNumber)
	fmt.Println(temp)
}
