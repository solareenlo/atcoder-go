package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	cnt, maxi := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			cnt++
			maxi = max(maxi, cnt)
		} else {
			cnt = 0
		}
	}

	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
