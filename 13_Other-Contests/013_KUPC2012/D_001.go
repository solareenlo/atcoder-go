package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	p := make([]int, 101)
	q := make([]int, 101)
	u := make([]int, 101)
	for i := 1; i <= m; i++ {
		fmt.Scan(&p[i], &q[i])
		for j := p[i]; j <= q[i]; j++ {
			u[j]++
		}
	}

	for i := 1; i <= n; i++ {
		if u[i] == 0 {
			fmt.Println("Impossible")
			return
		}
	}

	l := 0
	ans := 0
	for {
		res := 0
		for i := 1; i <= m; i++ {
			if p[i] <= l+1 {
				res = max(q[i], res)
			}
		}
		ans++
		l = res
		if res == n {
			break
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
