package main

import "fmt"

func main() {
	var n, a, b int
	var s string
	fmt.Scan(&n, &a, &b, &s)
	s += s
	c := 0
	ans := int(1e18)
	for i := 0; i < n; i++ {
		c = 0
		for j := 0; j < n/2; j++ {
			if s[i+j] != s[n+i-1-j] {
				c++
			}
		}
		ans = min(ans, c*b+a*i)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
