package main

import "fmt"

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)
	a = x / a
	b = y / b
	fmt.Println(min(a, b))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
