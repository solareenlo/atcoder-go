package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Print("A:")
	i := 0
	for i < a {
		fmt.Print("]")
		i++
	}
	fmt.Println()

	fmt.Print("B:")
	i = 0
	for i < b {
		fmt.Print("]")
		i++
	}
	fmt.Println()
}
