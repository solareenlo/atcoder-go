package main

import (
	"bufio"
	"fmt"
	"os"
)

const Nx = 200010

type edge struct {
	to, nex int
}

var a [2 * Nx]edge
var head [Nx]int
var cnt int

func add(u, v int) {
	cnt++
	a[cnt].to = v
	a[cnt].nex = head[u]
	head[u] = cnt
}

var N, Q int
var dep, siz, sd, sv [Nx]int
var fa [Nx][20]int

func dfs(p, f int) {
	dep[p] = dep[f] + 1
	siz[p] = 1
	fa[p][0] = f
	for i := 1; i <= 18; i++ {
		fa[p][i] = fa[fa[p][i-1]][i-1]
	}
	sd[p] = 0
	for i := head[p]; i > 0; i = a[i].nex {
		if a[i].to == f {
			continue
		}
		dfs(a[i].to, p)
		sd[p] += sd[a[i].to]
		siz[p] += siz[a[i].to]
	}
	sd[p] += siz[p] - 1
}

func dfs2(p, f int) {
	if p != 1 {
		sv[p] = sv[f] - siz[p] + (N - siz[p])
	}
	for i := head[p]; i > 0; i = a[i].nex {
		if a[i].to == f {
			continue
		}
		dfs2(a[i].to, p)
	}
}

func kthanc(x, k int) int {
	for i := 0; i <= 18; i++ {
		if (k & (1 << i)) != 0 {
			x = fa[x][i]
		}
	}
	return x
}

func LCA(x, y int) int {
	if dep[x] < dep[y] {
		x, y = y, x
	}
	w := dep[x] - dep[y]
	x = kthanc(x, w)
	if x == y {
		return x
	}
	for i := 18; i >= 0; i-- {
		if fa[x][i] != fa[y][i] {
			x = fa[x][i]
			y = fa[y][i]
		}
	}
	return fa[x][0]
}

func getans(x, y int) int {
	if x == y {
		return sv[x]
	}
	if dep[x] < dep[y] {
		x, y = y, x
	}
	var ret, md int
	lca := LCA(x, y)
	dis := dep[x] + dep[y] - 2*dep[lca]
	if dep[x] == dep[y] {
		md = kthanc(x, dis/2)
		sx := kthanc(x, dis/2-1)
		sy := kthanc(y, dis/2-1)
		ret = sv[md] + sv[x] + sv[y] - sv[sx] - sv[sy] - (dis/2)*N + (N-siz[sx]-siz[sy])*2
	} else {
		if dis%2 == 1 {
			md = kthanc(x, dis/2)
			ret = sv[x] + sv[y] - sv[md] - siz[md] - N*(dis/2)
		} else {
			md = kthanc(x, dis/2-1)
			ret = sv[x] + sv[y] - sv[md] - siz[md]*2 - N*(dis/2-1)
		}
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 1; i < +N; i++ {
		var j, k int
		fmt.Fscan(in, &j, &k)
		add(j, k)
		add(k, j)
	}
	dfs(1, 1)
	for i := 1; i <= N; i++ {
		sv[1] += dep[i]
	}
	sv[1] -= N
	dfs2(1, 1)
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var j, k int
		fmt.Fscan(in, &j, &k)
		fmt.Println(getans(j, k))
	}
}
