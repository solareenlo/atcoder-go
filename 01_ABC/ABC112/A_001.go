package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println("Hello World")
	} else {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
