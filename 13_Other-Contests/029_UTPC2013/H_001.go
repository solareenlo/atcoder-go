package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 1005
	const INF = 1000010000

	var n, m int
	fmt.Fscan(in, &n, &m)
	p := make([]int, MX)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	q := make([]int, MX)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
	}
	sv := 2 * n

	v := make([]int, 0)
	u := make([]int, 0)
	c := make([]int, 0)
	var add func(int, int, int)
	add = func(x, y, z int) {
		v = append(v, x)
		u = append(u, y)
		c = append(c, z)
	}
	for i := 0; i < n; i++ {
		add(sv, i, p[i])
		add(i, sv, 0)
		add(sv, n+i, 0)
		add(n+i, sv, q[i])
	}

	for i := 0; i < m; i++ {
		var x, y, a, b int
		fmt.Fscan(in, &x, &y, &a, &b)
		x--
		y--
		add(x, n+y, -a)
		add(n+y, x, b)
	}

	dist := make([]int, MX*2)
	for i := 0; i < n*2; i++ {
		dist[i] = INF
	}

	end := false
	for ti := 0; ti < n+5; ti++ {
		up := false
		for i := 0; i < len(v); i++ {
			if dist[v[i]] == INF {
				continue
			}
			if dist[u[i]] > dist[v[i]]+c[i] {
				dist[u[i]] = max(-INF, dist[v[i]]+c[i])
				up = true
			}
		}
		if !up {
			end = true
		}
		if !up || dist[sv] < 0 {
			break
		}
	}

	if end {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
