package main

import "fmt"

var (
	m int
	a = make([]int, 14*13*2)
	b = make([]int, 14*13*2)
	c = make([]int, 14*13*2)
)

func check(S int) bool {
	for i := 0; i < m; i++ {
		if S>>a[i]&1 != 0 && S>>b[i]&1 != 0 && S>>c[i]&1 != 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
		a[i]--
		b[i]--
		c[i]--
	}

	ans := 0
	for S := 0; S < 1<<n; S++ {
		if check(S) {
			cnt := 0
			for i := 0; i < n; i++ {
				if ((^S)>>i)&1 != 0 && !check(S|1<<i) {
					cnt++
				}
			}
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
