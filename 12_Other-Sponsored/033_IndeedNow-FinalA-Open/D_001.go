package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, flag int
}

var G [100001][]edge
var res int
var A, B [100001]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if c == 2 {
			G[a] = append(G[a], edge{b, 0})
			G[b] = append(G[b], edge{a, 0})
		} else {
			G[a] = append(G[a], edge{b, 1})
			G[b] = append(G[b], edge{a, 2})
		}
	}
	dfs(1, -1)
	fmt.Println(res - 1)
}

func dfs(pos, prev int) {
	for _, e := range G[pos] {
		if e.to == prev {
			continue
		}
		dfs(e.to, pos)
		if e.flag != 2 {
			res = max(res, A[e.to]+1+B[pos])
		}
		if e.flag != 1 {
			res = max(res, A[pos]+B[e.to]+1)
		}
		if e.flag != 2 {
			A[pos] = max(A[pos], A[e.to]+1)
		}
		if e.flag != 1 {
			B[pos] = max(B[pos], B[e.to]+1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
