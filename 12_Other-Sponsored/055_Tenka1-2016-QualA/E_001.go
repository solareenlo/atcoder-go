package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type edge struct {
	v, l, nxt int
}

const N = 300000
const INF = math.MaxInt - 1

var tot int
var e [N]edge
var head, vis [N]int
var cyc int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		add(x, y, 1)
		add(y, x, -1)
	}
	for i := range vis {
		vis[i] = INF
	}
	ans := 0
	for i := 0; i < n; i++ {
		if vis[i] == INF {
			cyc = 0
			DFS(i, 0)
			if cyc == 0 {
				fmt.Println(-1)
				return
			}
			ans += cyc
		}
	}
	fmt.Println(ans)
}

func add(x, y, z int) {
	tot++
	e[tot].v = y
	e[tot].l = z
	e[tot].nxt = head[x]
	head[x] = tot
}

func DFS(u, c int) {
	vis[u] = c
	for i, v := head[u], 0; i > 0; i = e[i].nxt {
		v = e[i].v
		if vis[v] == INF {
			DFS(v, c+e[i].l)
		} else {
			cyc = gcd(cyc, abs(c+e[i].l-vis[v]))
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
