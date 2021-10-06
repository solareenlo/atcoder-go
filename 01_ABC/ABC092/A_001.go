package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(min(a, b) + min(c, d))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
