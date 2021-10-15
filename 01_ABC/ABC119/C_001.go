package main

import "fmt"

var (
	N       int
	A, B, C int
	l       [10]int
	res     int
)

func dfs(p, a, b, c, val int) {
	if a != 0 && b != 0 && c != 0 {
		res = min(res, abs(a-A)+abs(b-B)+abs(c-C)+val-30)
	}
	if p > N {
		return
	}
	dfs(p+1, a+l[p], b, c, val+10)
	dfs(p+1, a, b+l[p], c, val+10)
	dfs(p+1, a, b, c+l[p], val+10)
	dfs(p+1, a, b, c, val)
}

func main() {
	fmt.Scan(&N, &A, &B, &C)
	for i := 1; i <= N; i++ {
		fmt.Scan(&l[i])
	}
	res = int(1e9)
	dfs(0, 0, 0, 0, 0)
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
