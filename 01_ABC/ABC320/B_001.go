package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			flag := true
			for k := 0; i+k <= j; k++ {
				if s[i+k] != s[j-k] {
					flag = false
					break
				}
			}
			if flag {
				ans = max(ans, j-i+1)
			}
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
