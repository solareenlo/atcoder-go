package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	x, y, z int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	e := make([][]int, N)
	ew := make([][]int, N)
	for i := 1; i <= m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		e[u] = append(e[u], v)
		ew[u] = append(ew[u], w)
		e[v] = append(e[v], u)
		ew[v] = append(ew[v], w)
	}

	flg := make([]bool, N)
	for i := 1; i <= k; i++ {
		var u int
		fmt.Fscan(in, &u)
		flg[u] = true
	}

	hd, tl := N<<2, N<<2
	tl++

	q := make([]node, N<<3)
	var vst [2][N]bool
	q[tl] = node{0, 1, 0}
	for hd < tl {
		hd++
		x := q[hd].x
		y := q[hd].y
		z := q[hd].z
		if vst[x][y] {
			continue
		}
		vst[x][y] = true
		if y == n {
			fmt.Println(z)
			return
		}
		if flg[y] && !vst[x^1][y] {
			q[hd] = node{x ^ 1, y, z}
			hd--
		}
		for i := 0; i < len(e[y]); i++ {
			v := e[y][i]
			w := ew[y][i]
			if (w ^ x) != 0 {
				if !vst[x][v] {
					tl++
					q[tl] = node{x, v, z + 1}
				}
			}
		}
	}
	fmt.Println(-1)
}
