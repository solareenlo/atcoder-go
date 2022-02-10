package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005
const mod = 1000000007

type edge struct{ v, nex int }

var (
	n    int
	ans  int
	e    = make([]edge, N<<1)
	head = make([]int, N)
	cnt  int
	fac  = make([]int, N)
	inv  = make([]int, N)
	X    int
	Y    int
	jud  int
	p    = make([]int, N)
	tot  int
	fat  = make([]int, N)
	pre  = make([]int, N)
	siz  = make([]int, N)
	g    = make([][]int, N)
	vis  = make([]bool, N)
)

func addedge(u, v int) {
	cnt++
	e[cnt].v = v
	e[cnt].nex = head[u]
	head[u] = cnt
}

func INIT() {
	fac[0] = 1
	inv[0] = 1
	inv[1] = 1
	for i := 1; i <= N-5; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	for i := 2; i <= N-5; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}
	ans = fac[n<<1]
}

func dfs(x, fa int) {
	vis[x] = true
	tot++
	p[tot] = x
	fat[x] = fa
	pre[x] = fa
	for i := head[x]; i > 0; i = e[i].nex {
		if e[i].v != fa {
			if !vis[e[i].v] {
				dfs(e[i].v, x)
			} else {
				jud++
				if X == 0 {
					X = e[i].v
				} else {
					Y = e[i].v
				}
			}
		}
	}
}

func dfs2(x int) {
	if siz[x] != 0 {
		return
	}
	siz[x] = 1
	for i := 0; i < len(g[x]); i++ {
		dfs2(g[x][i])
		siz[x] += siz[g[x][i]]
	}
}

func calc() int {
	for i := 1; i <= tot; i++ {
		siz[p[i]] = 0
		g[p[i]] = g[p[i]][:0]
	}
	for i := 1; i <= tot; i++ {
		if pre[pre[p[i]]] > p[i] {
			g[pre[p[i]]] = append(g[pre[p[i]]], p[i])
		}
	}
	for i := 1; i <= tot; i++ {
		dfs2(p[i])
	}
	rev := 1
	for i := 1; i <= tot; i++ {
		rev = rev * inv[siz[p[i]]] % mod
	}
	return rev
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n<<1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		addedge(x, y+n)
		addedge(y+n, x)
	}

	INIT()

	for i := 1; i <= n<<1; i++ {
		if !vis[i] {
			X = 0
			Y = 0
			jud = 0
			tot = 0
			dfs(i, 0)
			if jud > 2 {
				ans = 0
				break
			}
			for j := Y; j != i; j = fat[j] {
				pre[fat[j]] = j
			}
			pre[Y] = X
			ls := calc()
			for j := Y; j != X; j = fat[j] {
				pre[j] = fat[j]
			}
			pre[X] = Y
			ans = ans * (ls + calc()) % mod
		}
	}

	fmt.Println(ans)
}
