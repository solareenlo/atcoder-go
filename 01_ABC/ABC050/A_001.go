package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Scan(&a, &c, &b)

	switch c {
	case "+":
		fmt.Println(a + b)
	default:
		fmt.Println(a - b)
	}
}
