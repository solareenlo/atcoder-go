package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	if max(abs(r-8), abs(8-c))%2 != 0 {
		fmt.Println("black")
	} else {
		fmt.Println("white")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
