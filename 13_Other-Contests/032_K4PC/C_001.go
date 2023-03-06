package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e9)
	const N = 1 << 17

	type pair struct {
		x, y int
	}

	var n, q int
	fmt.Fscan(in, &n, &q)
	T := make([][]pair, N)
	for i := 1; i < n; i++ {
		var p, w int
		fmt.Fscan(in, &p, &w)
		p--
		T[p] = append(T[p], pair{i, w})
	}
	M := make([]int, n)
	for i := range M {
		M[i] = q
	}
	C := make([]int, q+1)
	for i := range C {
		C[i] = INF
	}
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		M[x] = i
	}
	var f func(int, int, int)
	f = func(u, c, m int) {
		if len(T[u]) == 0 {
			C[m] = min(C[m], c)
		}
		for _, tmp := range T[u] {
			v := tmp.x
			w := tmp.y
			f(v, c+w, min(m, M[v]))
		}
	}
	f(0, 0, q)
	for i := q; i >= 1; i-- {
		C[i-1] = min(C[i-1], C[i])
	}
	for i := 1; i <= q; i++ {
		if C[i] < INF {
			fmt.Println(C[i])
		} else {
			fmt.Println(-1)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
