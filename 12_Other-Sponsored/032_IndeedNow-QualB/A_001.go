package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(1 + abs(a-c) + abs(b-d))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
