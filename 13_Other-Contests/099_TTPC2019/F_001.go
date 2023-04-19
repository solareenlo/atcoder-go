package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 100005
	const INF = int(1e15)

	type P struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	var w, x, y, z int
	fmt.Fscan(in, &w, &x, &y, &z)
	w--
	x--
	y--
	z--
	var G [MX][]P
	for i := 0; i < m; i++ {
		var c, s, t int
		fmt.Fscan(in, &c, &s, &t)
		G[s-1] = append(G[s-1], P{t - 1, c})
	}
	w2 := make([]int, n)
	for i := range w2 {
		w2[i] = INF
	}
	y2 := make([]int, n)
	for i := range y2 {
		y2[i] = INF
	}
	wy2 := make([]int, n)
	for i := range wy2 {
		wy2[i] = INF
	}
	w2[w] = 0
	y2[y] = 0
	for i := 0; i < n; i++ {
		wy2[i] = min(wy2[i], w2[i]+y2[i])
		for _, ch := range G[i] {
			t := ch.x
			c := ch.y
			y2[t] = min(y2[t], y2[i]+c)
			w2[t] = min(w2[t], w2[i]+c)
			wy2[t] = min(wy2[t], wy2[i]+c)
		}
	}
	x2 := make([]int, n)
	for i := range x2 {
		x2[i] = INF
	}
	z2 := make([]int, n)
	for i := range z2 {
		z2[i] = INF
	}
	x2[x] = 0
	z2[z] = 0
	ans := w2[x] + y2[z]
	for i := n - 1; i >= 0; i-- {
		for _, ch := range G[i] {
			t := ch.x
			c := ch.y
			x2[i] = min(x2[i], x2[t]+c)
			z2[i] = min(z2[i], z2[t]+c)
		}
		ans = min(ans, x2[i]+z2[i]+wy2[i])
	}
	if ans >= INF {
		fmt.Println("Impossible")
	} else {
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
