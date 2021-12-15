package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, k int
	used = make([][]int, 200002)
	dp   = make([]int, 200002)
)

func dfs1(u, f int) {
	for _, v := range used[u] {
		if v == f {
			continue
		}
		dfs1(v, u)
		if dp[v] == 0 {
			dp[u]++
		}
	}
}

func dfs2(u, f int) {
	for _, v := range used[u] {
		if v == f {
			continue
		}
		tmp := -1
		if dp[v] != 0 {
			tmp = 0
		}
		if dp[u]+tmp == 0 {
			dp[v]++
		}
		dfs2(v, u)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		used[u] = append(used[u], v)
		used[v] = append(used[v], u)
	}

	dfs1(1, 1)
	dfs2(1, 1)

	res := 0
	for i := 1; i < n+1; i++ {
		if dp[i] == 0 {
			res++
		}
	}
	fmt.Println(res)
}
