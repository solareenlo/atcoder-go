package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := 0
	for i := 1; i <= n; i = i*10 + 1 {
		for j := 1; i*j <= n; j *= 10 {
			s += min(j, n-i*j+1)
		}
	}
	fmt.Println(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
