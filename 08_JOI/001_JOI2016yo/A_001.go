package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	A, B := 0, 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a > b {
			A += (a + b)
		}
		if a < b {
			B += (a + b)
		}
		if a == b {
			A += a
			B += b
		}
	}
	fmt.Println(A, B)
}
