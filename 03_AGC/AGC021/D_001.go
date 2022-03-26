package main

import "fmt"

var f = [301][301][301]int{}
var a string

func dp(l, r, k int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}
	if f[l][r][k] != 0 {
		return f[l][r][k]
	}
	f[l][r][k] = max(dp(l+1, r, k), dp(l, r-1, k))
	if a[l] == a[r] {
		f[l][r][k] = max(f[l][r][k], dp(l+1, r-1, k)+2)
	} else if k != 0 {
		f[l][r][k] = max(f[l][r][k], dp(l+1, r-1, k-1)+2)
	}
	return f[l][r][k]
}

func main() {
	var k int
	fmt.Scan(&a, &k)
	n := len(a)
	a = " " + a
	fmt.Println(dp(1, n, k))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
