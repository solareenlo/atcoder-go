package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	G := make([][]int, N)
	for i := 1; i < N; i++ {
		var P int
		fmt.Fscan(in, &P)
		P--
		G[P] = append(G[P], i)
	}
	memo := make([]int, N)
	var dfs func(int) int
	dfs = func(v int) int {
		res := 1
		for _, u := range G[v] {
			res += dfs(u)
		}
		memo[v] = res
		return memo[v]
	}
	dfs(0)
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, memo[i])
	}
}
