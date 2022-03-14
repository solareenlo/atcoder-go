package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	tmp := 0
	if a == 1 && b == 1 {
		tmp = 1
	}
	fmt.Println((max(4-a, 0) + max(4-b, 0) + tmp*4) * 100000)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
