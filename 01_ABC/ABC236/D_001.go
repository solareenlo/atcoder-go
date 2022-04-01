package main

import "fmt"

var (
	n   int
	ans int
	b   = make([]bool, 20)
	a   = [20][20]int{}
)

func dfs(t, m int) {
	if t > 2*n-1 {
		ans = max(ans, m)
		return
	}
	if b[t] {
		dfs(t+1, m)
		return
	}
	for j := t + 1; j <= 2*n; j++ {
		if !b[j] {
			if t == 0 {
				b[j] = false
			} else {
				b[j] = true
			}
			dfs(t+1, m^a[t][j])
			b[j] = false
		}
	}
}

func main() {
	fmt.Scan(&n)
	for i := 1; i < 2*n; i++ {
		for j := i + 1; j <= 2*n; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	dfs(1, 0)
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
