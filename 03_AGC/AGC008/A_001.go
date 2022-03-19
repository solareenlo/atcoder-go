package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	tmp := 2
	if x < y {
		tmp = 0
	}
	fmt.Println(min(abs(x+y)+1, abs(x-y)+tmp))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
