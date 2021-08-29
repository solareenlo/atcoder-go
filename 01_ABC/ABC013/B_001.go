package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a < b {
		a, b = b, a
	}

	if a-b > 5 {
		fmt.Println(10 - a + b)
	} else {
		fmt.Println(a - b)
	}
}
