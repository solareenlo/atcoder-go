package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	maxi, num := -10000, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			num++
			maxi = max(maxi, num)
		} else {
			num--
		}
	}
	fmt.Println(max(maxi, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
