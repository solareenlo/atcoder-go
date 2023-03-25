package main

import (
	"fmt"
)

const (
	N = 16
	P = 1e9 + 7
)

var (
	dp                  [N][1 << N]int
	matrix              [N][N]int
	vis                 [N]int
	n, a, b, m, ans, id int
)

func DFS(u int) int {
	ans := 1
	vis[u] = 1
	for i := 0; i < n; i++ {
		if matrix[u][i] != 0 && vis[i] == 0 && (((id>>u)&1)^((id>>i)&1)) != 0 {
			ans += DFS(i)
		}
	}
	return ans
}

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= m; i++ {
		fmt.Scan(&a, &b)
		a--
		b--
		matrix[a][b] = 1
		matrix[b][a] = 1
	}
	for i := 0; i < (1 << n); i++ {
		id = i
		for j := 0; j < n; j++ {
			vis[j] = 0
		}
		if DFS(0) == n {
			ans++
		}
	}
	fmt.Println(ans / 2)
}
