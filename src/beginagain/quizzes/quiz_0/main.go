package main

import (
	"fmt"
)

func main() {

	questions := [3][2]string{
		{"q: What is the most intelligent species of whale?", "cachalot"},
		{"q: What is my favourite colour?", "yellow"},
		{"q: What mammal is capable of powered flight?", "bats"},
	}

	for i := 0; i <= 3; i++ {
		fmt.Println(questions[i])
	}
}

reader := bufio.NewReader(os.Stdin)
fmt.Print("What is another intelligent species on Earth?> ")
input, _ := reader.ReadString('\n')
if strings.TrimRight(strings.ToLower(input), "\n") == "cachalot" {
	fmt.Println("Correct.")
} else {
	fmt.Println("What? Try again.")
}
}
