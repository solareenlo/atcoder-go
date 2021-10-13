package main

import "fmt"

var (
	p = make([]int, 51)
	b = make([]int, 51)
)

func dfs(n, x int) int {
	if n == 0 {
		if x <= 0 {
			return 0
		}
		return 1
	}
	if x <= 1+b[n-1] {
		return dfs(n-1, x-1)
	}
	return p[n-1] + 1 + dfs(n-1, x-2-b[n-1])
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	p[0], b[0] = 1, 1
	for i := 0; i < 50; i++ {
		b[i+1] = b[i]*2 + 3
		p[i+1] = p[i]*2 + 1
	}
	fmt.Println(dfs(n, x))
}
