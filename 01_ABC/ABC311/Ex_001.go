package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, X int
	fmt.Fscan(in, &N, &X)
	P := make([]int, N)
	for i := 1; i < N; i++ {
		var p int
		fmt.Fscan(in, &p)
		P[i] = p - 1
	}
	B := make([]int, N)
	W := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &B[i], &W[i], &C[i])
	}
	children := make([][]int, N)
	for i := 1; i < N; i++ {
		children[P[i]] = append(children[P[i]], i)
	}
	sz := make([]int, N)
	for i := range sz {
		sz[i] = 1
	}
	for i := N - 1; i >= 1; i-- {
		sz[P[i]] += sz[i]
	}
	for i := 0; i < N; i++ {
		sort.Slice(children[i], func(l, r int) bool {
			return sz[children[i][l]] > sz[children[i][r]]
		})
	}
	ans := make([]int, N)
	var dfs func(int, []int, bool) ([]int, []int)
	dfs = func(v int, dest []int, root bool) ([]int, []int) {
		if len(children[v]) == 0 {
			X0 := make([]int, len(dest))
			copy(X0, dest)
			X1 := make([]int, len(dest))
			copy(X1, dest)
			ans[v] = B[v]
			if C[v] == 1 {
				X0, X1 = X1, X0
			}
			for x := W[v]; x <= X; x++ {
				X0[x] = max(X0[x], X1[x-W[v]]+B[v])
			}
			if C[v] == 1 {
				X0, X1 = X1, X0
			}
			return X0, X1
		}
		X0, X1 := dfs(children[v][0], dest, root)
		for i := 1; i < len(children[v]); i++ {
			X0, _ = dfs(children[v][i], X0, false)
			_, X1 = dfs(children[v][i], X1, false)
		}
		if C[v] == 1 {
			X0, X1 = X1, X0
		}
		if root {
			for x := W[v]; x <= X; x++ {
				ans[v] = max(ans[v], X1[x-W[v]]+B[v])
			}
		}
		for x := W[v]; x <= X; x++ {
			X0[x] = max(X0[x], X1[x-W[v]]+B[v])
		}
		if C[v] == 1 {
			X0, X1 = X1, X0
		}
		return X0, X1
	}
	for i := 0; i < N; i++ {
		if i == 0 || children[P[i]][0] != i {
			tmp := make([]int, X+1)
			dfs(i, tmp, true)
		}
	}
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
