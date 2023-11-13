package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(max(x+z, y))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
