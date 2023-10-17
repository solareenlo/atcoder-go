package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(abs(a) + abs(a-b) + abs(b))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
