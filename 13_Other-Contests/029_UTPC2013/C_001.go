package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 1010

var dep [2]int
var e [][][]int

func bfs(s, ii int) int {
	q := make([]int, 0)
	dep := make([]int, maxn)
	for i := range dep {
		dep[i] = -1
	}
	dep[s] = 0
	maxx := 0
	q = append(q, s)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		maxx = dep[u]
		for i := 0; i < len(e[u][ii]); i++ {
			v := e[u][ii][i]
			if dep[v] == -1 {
				dep[v] = dep[u] + 1
				q = append(q, v)
			}
		}
	}
	return maxx
}

func main() {
	in := bufio.NewReader(os.Stdin)

	e = make([][][]int, maxn)
	for i := range e {
		e[i] = make([][]int, 2)
		for j := range e[i] {
			e[i][j] = make([]int, 0)
		}
	}
	d := make([]int, 2)
	ansmx, ansmi := 1, 1
	for j := 0; j < 2; j++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		for i := 1; i <= m; i++ {
			var u, v int
			fmt.Fscan(in, &u, &v)
			e[u][j] = append(e[u][j], v)
			e[v][j] = append(e[v][j], u)
		}
		mx := -1145141919
		mi := 2147483646
		for i := 0; i < n; i++ {
			t := bfs(i, j)
			mx = max(mx, t)
			mi = min(mi, t)
		}
		d[j] = mx
		ansmx += mx
		ansmi += mi
	}
	fmt.Println(max(max(d[0], d[1]), ansmi), ansmx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
