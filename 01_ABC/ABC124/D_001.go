package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)
	s = "2" + s

	m := make(map[int]int)
	cnt, res := 0, 0
	for i := 1; i <= n; i++ {
		if s[i] == '0' && s[i] != s[i-1] {
			cnt++
		}
		if s[i] == '0' {
			m[cnt] = i
		}
		res = max(res, i-m[cnt-k])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
