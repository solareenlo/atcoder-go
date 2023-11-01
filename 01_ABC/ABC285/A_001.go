package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if max(a, b)/2 == min(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
