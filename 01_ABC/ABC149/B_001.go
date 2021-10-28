package main

import "fmt"

func main() {
	var A, B, k int
	fmt.Scan(&A, &B, &k)

	var a, b int
	if k >= A {
		a = 0
	} else {
		a = A - k
	}
	if k-A > 0 {
		b = max(0, B-(k-A))
	} else {
		b = B
	}

	fmt.Println(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
