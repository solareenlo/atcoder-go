package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if abs(x-y) < 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
