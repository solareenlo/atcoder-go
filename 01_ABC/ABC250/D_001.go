package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	p := make([]int, 1<<20)
	t := make([]int, 1<<20)
	for i := 2; i*i*i <= n; i++ {
		t[i] = t[i-1]
		if p[i] == 0 {
			t[i]++
			ans += t[min(n/i/i/i, i-1)]
			for j := i + i; j < 1<<20; j += i {
				p[j] = 1
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
