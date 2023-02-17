package main

import (
	"bufio"
	"fmt"
	"os"
)

const mN = 1010
const mM = 100010

var n, m, oe int
var head [mN]int
var ver, nxt [4 * mM]int
var vis []bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	if m+n-1 <= n*(n-1)/2 {
		for i := 1; i <= m; i++ {
			var x int
			fmt.Fscan(in, &x)
			fmt.Println("no")
		}
		return
	}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			modify(i, j)
			modify(j, i)
		}
	}
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		modify(x, y)
		modify(y, x)
		if check() {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func modify(x, y int) {
	if ver[head[x]] == y {
		head[x] = nxt[head[x]]
		return
	}
	for i := head[x]; nxt[i] > 0; i = nxt[i] {
		if ver[nxt[i]] == y {
			nxt[i] = nxt[nxt[i]]
			return
		}
	}
	oe++
	nxt[oe] = head[x]
	ver[oe] = y
	head[x] = oe
}

func check() bool {
	vis = make([]bool, mN)
	for i := 1; i <= n; i++ {
		if !vis[i] && !dfs(i, 0) {
			return false
		}
	}
	return true
}

func dfs(x, f int) bool {
	vis[x] = true
	for t := head[x]; t > 0; t = nxt[t] {
		y := ver[t]
		if y == f {
			continue
		}
		if vis[y] || !dfs(y, x) {
			return false
		}
	}
	return true
}
