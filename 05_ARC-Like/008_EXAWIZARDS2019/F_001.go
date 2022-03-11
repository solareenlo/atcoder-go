package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct{ x, y, d int }

const N = 100005

var (
	n  int
	m  int
	R1 = make([]node, 0)
	R2 = make([]node, 0)
	C1 = make([]node, 0)
	C2 = make([]node, 0)
	x  = [N][2]int{}
	y  = [N][2]int{}
	r  = [N]int{}
	c  = [N]int{}
)

func dfs1R(k, a, b, d int) {
	if a < 1 || a > n || b < 1 || b > m {
		return
	}
	R1 = append(R1, node{a, b, d})
	if k == 0 {
		return
	}
	dfs1C(k-1, a, b, d)
	dfs1C(k-1, a, y[b][r[a]], d+abs(y[b][r[a]]-b))
}

func dfs1C(k, a, b, d int) {
	if a < 1 || a > n || b < 1 || b > m {
		return
	}
	C1 = append(C1, node{a, b, d})
	if k == 0 {
		return
	}
	dfs1R(k-1, a, b, d)
	dfs1R(k-1, x[a][c[b]], b, d+abs(x[a][c[b]]-a))
}

func dfs2R(k, a, b, d int) {
	if a < 1 || a > n || b < 1 || b > m {
		return
	}
	R2 = append(R2, node{a, b, d})
	if k == 0 {
		return
	}
	dfs2C(k-1, a, b, d)
	dfs2C(k-1, a, y[b][1-r[a]], d+abs(y[b][1-r[a]]-b))
}

func dfs2C(k, a, b, d int) {
	if a < 1 || a > n || b < 1 || b > m {
		return
	}
	C2 = append(C2, node{a, b, d})
	if k == 0 {
		return
	}
	dfs2R(k-1, a, b, d)
	dfs2R(k-1, x[a][1-c[b]], b, d+abs(x[a][1-c[b]]-a))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &n, &m, &Q)
	var sr, sc string
	fmt.Fscan(in, &sr, &sc)
	sr = " " + sr
	sc = " " + sc

	r[0] = -1
	r[n+1] = -1
	c[0] = -1
	c[m+1] = -1
	for i := 1; i < n+1; i++ {
		if sr[i] == 'E' {
			r[i] = 1
		}
	}
	for i := 1; i < m+1; i++ {
		if sc[i] == 'S' {
			c[i] = 1
		}
	}
	for i := 1; i < n+1; i++ {
		if r[i] != r[i-1] {
			x[i][0] = i - 1
		} else {
			x[i][0] = x[i-1][0]
		}
	}
	for i := n; i >= 1; i-- {
		if r[i] != r[i+1] {
			x[i][1] = i + 1
		} else {
			x[i][1] = x[i+1][1]
		}
	}
	for i := 1; i < m+1; i++ {
		if c[i] != c[i-1] {
			y[i][0] = i - 1
		} else {
			y[i][0] = y[i-1][0]
		}
	}
	for i := m; i >= 1; i-- {
		if c[i] != c[i+1] {
			y[i][1] = i + 1
		} else {
			y[i][1] = y[i+1][1]
		}
	}

	dir := [2]int{-1, 1}
	for i := 0; i < Q; i++ {
		var a, b, u, v int
		fmt.Fscan(in, &a, &b, &u, &v)
		R1 = R1[:0]
		C1 = C1[:0]
		R2 = R2[:0]
		C2 = C2[:0]
		dfs1R(2, a, b, 0)
		dfs1C(2, a, b, 0)
		dfs2R(2, u, v, 0)
		dfs2C(2, u, v, 0)
		ans := 1 << 60
		for _, p := range R1 {
			for _, q := range C2 {
				R := (q.y - p.y) * dir[r[p.x]]
				C := (q.x - p.x) * dir[c[q.y]]
				if R >= 0 && C >= 0 {
					ans = min(ans, R+C+p.d+q.d)
				}
			}
		}
		for _, p := range C1 {
			for _, q := range R2 {
				R := (q.y - p.y) * dir[r[q.x]]
				C := (q.x - p.x) * dir[c[p.y]]
				if R >= 0 && C >= 0 {
					ans = min(ans, R+C+p.d+q.d)
				}
			}
		}
		if ans >= 1<<60 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, ans)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
