package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
