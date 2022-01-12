package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if abs(a) == abs(b) {
		fmt.Println("Draw")
	} else if abs(a) < abs(b) {
		fmt.Println("Ant")
	} else {
		fmt.Println("Bug")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
