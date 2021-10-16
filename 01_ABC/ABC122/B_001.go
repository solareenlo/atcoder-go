package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	maxi, cnt := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' || s[i] == 'C' || s[i] == 'G' || s[i] == 'T' {
			cnt++
		} else {
			cnt = 0
		}
		maxi = max(maxi, cnt)
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
