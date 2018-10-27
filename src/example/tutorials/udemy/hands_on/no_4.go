package main

import "fmt"

func main() {
	m := map[string]string{"thing": "montana", "otherthing": "deer"}
	fmt.Println(m)

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
