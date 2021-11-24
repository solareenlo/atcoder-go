package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	edge := make([][]int, n)
	weight := make([][]int, n)
	for i := 1; i < n; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		edge[u] = append(edge[u], v)
		edge[v] = append(edge[v], u)
		weight[u] = append(weight[u], w)
		weight[v] = append(weight[v], w)
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0

	que := make([]int, 0)
	que = append(que, 0)
	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		for i := 0; i < len(edge[now]); i++ {
			next := edge[now][i]
			sum := dist[now] ^ weight[now][i]
			if dist[next] == -1 {
				dist[next] = sum
				que = append(que, next)
			}
		}
	}

	mod := int(1e9 + 7)
	res := 0
	for i := 0; i < 60; i++ {
		cnt := make([]int, 2)
		for j := 0; j < n; j++ {
			cnt[dist[j]>>i&1]++
		}
		res += (1 << i) % mod * cnt[0] % mod * cnt[1]
		res %= mod
	}
	fmt.Println(res)
}
