package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005
const M = 800040

var (
	n    int
	A    int
	B    int
	d    int
	g    = [N]int{}
	ans  int
	a        = make([]arr, N)
	Head     = [N]int{}
	Next     = [M]int{}
	Adj      = [M]int{}
	tot  int = 0
	vis      = [N]bool{}
)

type arr struct{ x, y, id int }

func addedge(u, v int) {
	tot++
	Next[tot] = Head[u]
	Head[u] = tot
	Adj[tot] = v
	tot++
	Next[tot] = Head[v]
	Head[v] = tot
	Adj[tot] = u
}

func work(k int) {
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
	for i := 0; i < n; i++ {
		a[i+1] = tmp[i]
	}
	var i, j, l, r int
	for i, j, l, r = 1, 1, 1, 1; i <= n; i++ {
		for ; a[l].x < a[i].x-d || a[l].x == a[i].x-d && a[l].y < a[i].y-k; l++ {
		}
		for ; a[r].x < a[i].x-d || a[r].x == a[i].x-d && a[r].y <= a[i].y+k; r++ {
		}
		g[a[i].id] += r - l
		j = max(j, l)
		if j >= r {
			continue
		}
		for addedge(a[i].id, a[j].id); j+1 < r; j++ {
			addedge(a[j].id, a[j+1].id)
		}
	}
}

func dfs(x int) {
	vis[x] = true
	ans += g[x]
	for e := Head[x]; e > 0; e = Next[e] {
		if !vis[Adj[e]] {
			dfs(Adj[e])
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &A, &B)
	a = make([]arr, n+2)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[i] = arr{x + y, x - y, i}
	}
	d = max(abs(a[A].x-a[B].x), abs(a[A].y-a[B].y))
	work(d)
	for i := 1; i <= n; i++ {
		a[i].x, a[i].y = a[i].y, a[i].x
	}
	work(d - 1)
	dfs(A)
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
