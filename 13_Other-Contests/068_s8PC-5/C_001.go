package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)

	for j := 0; j < q; j++ {
		var s string
		fmt.Scan(&s)
		n := len(s)
		mx, mn, cur := 0, 0, 0
		for i := 0; i < n; i++ {
			if s[i] == '(' {
				cur++
			} else {
				cur--
			}
			mx = max(mx, cur)
			mn = min(mn, cur)
		}
		if cur == 0 && mx <= n/4 && mn >= -n/4 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
