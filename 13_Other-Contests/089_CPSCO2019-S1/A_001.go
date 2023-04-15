package main

import "fmt"

func main() {
	var x, a int
	fmt.Scan(&x, &a)
	fmt.Println((a+2)/3, min(a, x/3))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
