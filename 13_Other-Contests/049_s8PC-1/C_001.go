package main

import "fmt"

var D [50]int
var S [50]int
var G [50][]int
var ans int64

func dfs(u int, t int64) {
	t += int64(D[u])
	S[u] = 1
	if t > ans {
		ans = t
	}
	for _, v := range G[u] {
		if S[v] == 0 {
			dfs(v, t)
		}
	}
	S[u] = 0
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&D[i])
	}
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		x--
		y--
		G[x] = append(G[x], y)
		G[y] = append(G[y], x)
	}
	for i := 0; i < n; i++ {
		dfs(i, 0)
	}
	fmt.Println(ans)
}
