package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt1, cnt2 := 0, 0
	for i := 0; i < len(s); i++ {
		if (i%2 == 0 && s[i] == '1') || (i%2 == 1 && s[i] == '0') {
			cnt1++
		}
		if (i%2 == 0 && s[i] == '0') || (i%2 == 1 && s[i] == '1') {
			cnt2++
		}
	}

	fmt.Println(min(cnt1, cnt2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
