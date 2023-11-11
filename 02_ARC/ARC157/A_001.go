package main

import "fmt"

func main() {
	var n, a, b, c, d int
	fmt.Scan(&n, &a, &b, &c, &d)
	if abs(b-c) >= 2 || a != 0 && b == 0 && c == 0 && d != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
