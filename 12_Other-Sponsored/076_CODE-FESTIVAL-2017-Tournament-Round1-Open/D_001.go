package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 200200

type edge struct {
	r, nxt, w int
}

var e [maxn << 1]edge
var head [maxn]int
var esz, n, id int
var anses, s [maxn]int
var sum int

func addedge(u, v, w int) {
	esz++
	e[esz].r = v
	e[esz].nxt = head[u]
	head[u] = esz
	e[esz].w = w
	esz++
	e[esz].r = u
	e[esz].nxt = head[v]
	head[v] = esz
	e[esz].w = w
}

func dfs(u, f int) int {
	sz, xs := 1, 0
	for t := head[u]; t > 0; t = e[t].nxt {
		if e[t].r != f {
			xs = dfs(e[t].r, u)
			sz += xs
			if 2*xs != n {
				anses[e[t].w] = (s[e[t].r] - s[u]) / (n - 2*xs)
				sum -= anses[e[t].w] * xs * (n - xs) * 2
			} else {
				id = e[t].w
			}
		}
	}
	return sz
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 2; i <= n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		addedge(u, v, i)
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		sum += s[i]
	}
	dfs(1, 0)
	if id != 0 {
		anses[id] = sum * 2 / n / n
	}
	for i := 2; i <= n; i++ {
		fmt.Fprintln(out, anses[i])
	}
}
