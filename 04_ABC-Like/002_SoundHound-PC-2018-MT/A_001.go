package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a+b == 15 {
		fmt.Print("+")
	} else if a*b == 15 {
		fmt.Print("*")
	} else {
		fmt.Print("x")
	}
	fmt.Println()
}
