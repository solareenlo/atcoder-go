package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b%a != 0 {
		fmt.Println(b - a)
	} else {
		fmt.Println(a + b)
	}
}
