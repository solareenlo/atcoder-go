package main

import "fmt"

func main() {
	var a, b int
	var c, d string
	fmt.Scan(&c, &a, &d, &b)
	if c == d {
		fmt.Println(abs(a-b) / 15)
	} else {
		fmt.Println((a + b) / 15)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
