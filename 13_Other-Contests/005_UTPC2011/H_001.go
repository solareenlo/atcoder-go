package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	to, cap, cost, rev int
}

const N = 10100

var V int
var G [][]edge
var cost, prevv, preve []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n, k int
	fmt.Fscan(in, &m, &n, &k)
	V = k
	w := make([]int, N)
	a := make([]int, N)
	num := make([][]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}

	G = make([][]edge, N)
	prevv = make([]int, N)
	preve = make([]int, N)
	sum := 0
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		if i == 0 || a[i-1] != a[i] {
			sum += w[a[i]]
		}
		if i > 0 {
			addEdge(i-1, i, N, 0)
		}
		num[a[i]] = append(num[a[i]], i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(num[i])-1; j++ {
			if num[i][j] < num[i][j+1]-1 {
				addEdge(num[i][j], num[i][j+1]-1, 1, -w[i])
			}
		}
	}
	fmt.Println(sum + mcf(m-1))
}

func addEdge(from, to, cap, cost int) {
	G[from] = append(G[from], edge{to, cap, cost, len(G[to])})
	G[to] = append(G[to], edge{from, 0, -cost, len(G[from]) - 1})
}

func mcf(f int) int {
	res := 0
	for f > 0 {
		cost = make([]int, V)
		for i := 0; i < V; i++ {
			prevv[i] = i - 1
			if i == 0 {
				preve[i] = -1
				continue
			}
			for j := 0; j < len(G[i-1]); j++ {
				if G[i-1][j].to == i {
					preve[i] = j
					break
				}
			}
		}
		now := 0
		for now < V {
			next := now + 1
			for i := 0; i < len(G[now]); i++ {
				e := G[now][i]
				if e.cap > 0 && cost[e.to] > cost[now]+e.cost {
					cost[e.to] = cost[now] + e.cost
					prevv[e.to] = now
					preve[e.to] = i
					next = min(next, e.to)
				}
			}
			now = next
		}
		if cost[V-1] == 0 {
			break
		}
		fl := f
		for v := V - 1; v > 0; v = prevv[v] {
			fl = min(fl, G[prevv[v]][preve[v]].cap)
		}
		f -= fl
		res += fl * cost[V-1]
		for v := V - 1; v > 0; v = prevv[v] {
			e := G[prevv[v]][preve[v]]
			G[prevv[v]][preve[v]].cap -= fl
			G[v][e.rev].cap += fl
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
