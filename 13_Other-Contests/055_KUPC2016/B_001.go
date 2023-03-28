package main

import (
	"fmt"
)

var (
	n, k int
	a    [26]int
)

func main() {
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		var p string
		fmt.Scan(&p)
		a[p[0]-'A']++
	}
	l, r := -1, int(1e16)
	for r-l > 1 {
		m := (l + r) / 2
		cnt := (0)
		for i := 0; i < 26; i++ {
			cnt += min(m, a[i])
		}
		if cnt >= k*m {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
