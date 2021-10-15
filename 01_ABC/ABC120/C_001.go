package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt0, cnt1 := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
	}
	fmt.Println(min(cnt0, cnt1) * 2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
