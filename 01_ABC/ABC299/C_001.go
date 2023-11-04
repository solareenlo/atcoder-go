package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	cnt := 0
	ans := -1
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			cnt++
		} else {
			if cnt > 0 {
				ans = max(ans, cnt)
			}
			cnt = 0
		}
	}
	if cnt > 0 {
		if n != cnt {
			ans = max(ans, cnt)
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
