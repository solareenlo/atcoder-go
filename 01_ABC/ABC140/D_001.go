package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	cnt := 0
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		}
	}

	fmt.Println(min(n-1, cnt+2*k))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
