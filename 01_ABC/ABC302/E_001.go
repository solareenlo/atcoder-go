package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300300

type EDGE struct {
	v, w, nxt int
}

var edge [N * 2]EDGE
var cnt, ans int
var head, du [N]int

func add(u, v int) {
	edge[cnt].v = v
	edge[cnt].w = 1
	edge[cnt].nxt = head[u]
	head[u] = cnt
	cnt++
}

func Remove(u int) {
	if du[u] != 0 {
		ans++
	}
	du[u] = 0
	for i := head[u]; i != -1; i = edge[i].nxt {
		if edge[i].w != 0 {
			edge[i^1].w = 0
			du[edge[i].v]--
			if du[edge[i].v] == 0 {
				ans++
			}
		}
	}
	head[u] = -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans = n
	for i := 1; i <= n; i++ {
		head[i] = -1
	}
	for i := 1; i <= m; i++ {
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var u1, v1 int
			fmt.Fscan(in, &u1, &v1)
			add(u1, v1)
			add(v1, u1)
			if du[u1] == 0 {
				ans--
			}
			if du[v1] == 0 {
				ans--
			}
			du[u1]++
			du[v1]++
		}
		if op == 2 {
			var u1 int
			fmt.Fscan(in, &u1)
			Remove(u1)
		}
		fmt.Fprintln(out, ans)
	}
}
