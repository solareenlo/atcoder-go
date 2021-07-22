package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	switch {
	case a > b:
		fmt.Println("GREATER")
	case a < b:
		fmt.Println("LESS")
	default:
		fmt.Println("EQUAL")
	}
}
