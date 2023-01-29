package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(c%3*a + c/3*min(3*a, b))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
