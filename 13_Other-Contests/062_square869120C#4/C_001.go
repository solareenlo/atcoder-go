package main

import (
	"fmt"
)

var (
	t          [7][200020]bool
	vis        [7][200020]bool
	n, m, q, a int64
)

func DFS(x, y, lim int, vis *[7][200020]bool) {
	(*vis)[x][y] = true
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, 1, -1, 0}
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 0 && b >= 0 && a < 7 && b < lim {
			if !(*vis)[a][b] {
				DFS(a, b, lim, vis)
			}
		}
	}
}

func Sol(lim int64, t *[7][200020]bool) int64 {
	var ret int64 = 0
	for i := 0; i < 7; i++ {
		for j := int64(0); j < lim; j++ {
			vis[i][j] = t[i][j]
		}
	}
	for i := 0; i < 7; i++ {
		for j := int64(0); j < lim; j++ {
			if !vis[i][j] {
				ret++
				DFS(i, int(j), int(lim), &vis)
			}
		}
	}
	return ret
}

func main() {
	fmt.Scan(&n, &m, &q)
	for i := 0; i < 7; i++ {
		for j := int64(0); j < m; j++ {
			t[i][j] = false
		}
	}
	for q > 0 {
		fmt.Scan(&a)
		for i := a; i < 14*m; i += m {
			t[i%7][i/7] = true
		}
		q--
	}
	k := Sol(m, &t)
	w := Sol(m*2, &t)
	var ans int64
	g := n/m - 1
	if g <= 0 {
		ans = Sol(n, &t)
	} else {
		ans = (g)*(w-k) + Sol(n-(g*m), &t)
	}
	fmt.Println(ans)
}
