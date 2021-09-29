package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	if abs(x-a) > abs(x-b) {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
