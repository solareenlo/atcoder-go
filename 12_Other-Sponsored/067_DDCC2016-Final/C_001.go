package main

import "fmt"

func main() {
	var a, b, c int
	var s string
	fmt.Scan(&a, &b, &c, &s)
	n := len(s)
	v := make([]int, n+2)
	v[1] = int(s[0] - '0')
	for i := 1; i < n; i++ {
		v[i+1] = v[i]
		if s[i] != s[i-1] {
			v[i+1]++
		}
	}
	v[n+1] = v[n] + int('1'-s[n-1])
	ans := int(1e18)
	for i := 0; i <= n; i++ {
		ans = min(ans, a*i+b*(n-i)+max(v[i], v[n+1]-v[i+1])*c)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
