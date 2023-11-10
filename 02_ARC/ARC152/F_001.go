package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAXN = 200005

var fa, dep, siz [MAXN]int
var vis [MAXN]bool
var edg [MAXN][]int
var ans int

func dfs(u, f int) {
	fa[u] = f
	siz[u] = 1
	for _, v := range edg[u] {
		if v != f {
			dep[v] = dep[u] + 1
			dfs(v, u)
			siz[u] += siz[v]
		}
	}
}

func upd(u, lim int) {
	ans++
	for _, v := range edg[u] {
		if !vis[v] && siz[u] > lim {
			vis[v] = true
			siz[u] -= siz[v]
			upd(v, lim)
		}
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		edg[u] = append(edg[u], v)
		edg[v] = append(edg[v], u)
	}
	dfs(1, 0)
	if dep[n]%2 != n%2 {
		fmt.Println(-1)
		return
	}
	for i := 1; i <= n; i++ {
		sort.Slice(edg[i], func(x, y int) bool {
			return siz[edg[i][x]] > siz[edg[i][y]]
		})
	}
	pth := make([]int, 0)
	for i := n; i > 0; i = fa[i] {
		vis[i] = true
		pth = append(pth, i)
	}
	reverseOrderInt(pth)
	for _, i := range pth {
		siz[fa[i]] -= siz[i]
	}
	for i := n; i > 0; i = fa[i] {
		upd(i, (n+2*dep[i]-dep[n])/2)
	}
	fmt.Println(ans)
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
