package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)

	ans := 10000
	for i := 0; i < 10000; i++ {
		for j := 0; i+j < 10000; j++ {
			now := a*i + b*j
			if n >= now && (n-now)%c == 0 {
				ans = min(ans, i+j+(n-now)/c)
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
