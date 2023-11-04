package main

import "fmt"

func main() {
	var b [300300]int
	var n, m, k int
	var s string
	fmt.Scan(&n, &m, &k, &s)
	for i := 1; i <= n; i++ {
		if s[i-1] == 'x' {
			b[i] = b[i-1] + 1
		} else {
			b[i] = b[i-1]
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		l := i
		r := n * m
		for l <= r {
			mid := (l + r) / 2
			if b[n]*(mid/n)+b[mid%n]-b[i-1] <= k {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		ans = max(ans, r-i+1)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
