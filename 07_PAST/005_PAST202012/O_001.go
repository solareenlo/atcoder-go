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

	var N, M int
	fmt.Fscan(in, &N, &M)
	G := make([][]int, 2<<17)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	B := 500
	H := make([][]int, 2<<17)
	for i := 1; i <= N; i++ {
		if len(G[i]) >= B {
			for _, u := range G[i] {
				H[u] = append(H[u], i)
			}
			G[i] = G[i][:0]
		}
	}

	pos := make([]int, 2<<17)
	now := make([]int, 2<<17)
	pre := make([]int, 2<<17)
	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var T, x int
		fmt.Fscan(in, &T, &x)
		if T == 1 {
			pos[x]++
			for _, u := range G[x] {
				now[u]++
			}
		} else {
			ans := now[x]
			for _, u := range H[x] {
				ans += pos[u]
			}
			fmt.Fprintln(out, ans-pre[x])
			pre[x] = ans
		}
	}
}
