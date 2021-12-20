package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)
	n := len(s)

	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			cnt[i+1] = cnt[i] + 1
		} else {
			cnt[i+1] = cnt[i]
		}
	}

	res := 0
	r := 0
	for l := 0; l < n; l++ {
		for r <= n-1 && cnt[r+1]-cnt[l] <= k {
			r = r + 1
		}
		res = max(res, r-l)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
