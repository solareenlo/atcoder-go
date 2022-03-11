package main

import "fmt"

func main() {
	var K, A, B int
	fmt.Scan(&K, &A, &B)

	fmt.Println(max((K+1-A)/2*(B-A)+A+(K+1-A)%2, K+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
