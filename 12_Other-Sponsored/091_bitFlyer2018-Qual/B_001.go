package main

import "fmt"

func main() {
	var A, B, N int
	var X string
	fmt.Scan(&A, &B, &N, &X)
	for _, c := range X {
		if c == 'S' || (c == 'E' && A >= B) {
			A = max(0, A-1)
		} else {
			B = max(0, B-1)
		}
	}
	fmt.Println(A)
	fmt.Println(B)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
