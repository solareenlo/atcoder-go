package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if abs(x-y) > 1 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Brown")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
