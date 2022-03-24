package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct{ id, v int }

const N = 200005

var (
	n    int
	cntE int
	e    = make([][]Edge, N)
	dg   = make([]int, N)
	cur  = make([]int, N)
	vs   = make([]bool, 2*N)
	ans  = make([]int, N)
)

func addE(u, v int) {
	cntE++
	e[u] = append(e[u], Edge{cntE, v})
	dg[u]++
	e[v] = append(e[v], Edge{cntE, u})
	dg[v]++
}

func dfs(u int) {
	for i := cur[u]; i < len(e[u]); i = cur[u] {
		cur[u]++
		id := e[u][i].id
		v := e[u][i].v
		if vs[id] {
			continue
		}
		vs[id] = true
		dfs(v)
		if v == n+u {
			ans[u] = 1
		} else if u == n+v {
			ans[v] = -1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	S := n*2 + 1
	for i := 1; i <= n; i++ {
		var f int
		fmt.Fscan(in, &f)
		if f >= 0 {
			addE(i, f)
		} else {
			addE(i, S)
		}
	}

	for i := 1; i <= n; i++ {
		var f int
		fmt.Fscan(in, &f)
		if f >= 0 {
			addE(n+i, n+f)
		} else {
			addE(n+i, S)
		}
	}

	for i := 1; i <= n; i++ {
		if (dg[i]&1)^(dg[n+i]&1) != 0 {
			fmt.Println("IMPOSSIBLE")
			return
		} else if dg[i]&1 != 0 {
			addE(i, n+i)
		}
	}

	fmt.Println("POSSIBLE")
	dfs(1)
	for i := 1; i <= n; i++ {
		fmt.Print(ans[i], " ")
	}
}
