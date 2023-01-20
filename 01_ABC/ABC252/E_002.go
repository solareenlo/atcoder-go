package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type ab struct {
	a, to, cost, next int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	tot := 1
	N := 200005
	edge := make([]ab, N*2)
	h := make([]int, N)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		tot++
		edge[tot].a = a
		edge[tot].to = b
		edge[tot].cost = c
		edge[tot].next = h[a]
		h[a] = tot
		tot++
		edge[tot].a = b
		edge[tot].to = a
		edge[tot].cost = c
		edge[tot].next = h[b]
		h[b] = tot
	}

	dist := make([]int, N)
	for i, _ := range dist {
		dist[i] = math.MaxInt64
	}

	p := make([]int, N)
	q := make([]int, N*2)
	q[1] = 1
	dist[1] = 0
	j := 1
	for i := 1; i <= j; i++ {
		a := q[i]
		for k := h[a]; k > 0; k = edge[k].next {
			to := edge[k].to
			cost := edge[k].cost
			if dist[a]+cost < dist[to] {
				dist[to] = dist[a] + cost
				j++
				q[j] = to
				p[to] = k >> 1
			}
		}
	}
	for i := 2; i <= n; i++ {
		fmt.Fprintf(out, "%d ", p[i])
	}
}
