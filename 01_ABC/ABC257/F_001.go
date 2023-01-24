package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 300005

var tot int
var head [MAXN]int
var to [MAXN * 2]int
var nxt [MAXN * 2]int
var n, m int
var q [MAXN]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)

	var x, y int
	for m > 0 {
		fmt.Fscan(in, &x, &y)
		add(x, y)
		add(y, x)
		m--
	}

	dis1 := make([]int, MAXN)
	disn := make([]int, MAXN)
	bfs(dis1, 1)
	bfs(disn, n)
	for i := 1; i <= n; i++ {
		ans := min(dis1[n], dis1[i]+disn[0], dis1[0]+disn[i])
		if ans >= 1e9 {
			fmt.Fprint(out, "-1 ")
		} else {
			fmt.Fprint(out, ans, " ")
		}
	}
}

func bfs(dis []int, s int) {
	for i := 0; i <= n; i++ {
		dis[i] = 1e9
	}
	L, R := 1, 1
	q[1] = s
	dis[s] = 0
	for L <= R {
		p := q[L]
		L++
		for i := head[p]; i > 0; i = nxt[i] {
			if dis[to[i]] > dis[p]+1 {
				dis[to[i]] = dis[p] + 1
				R++
				q[R] = to[i]
			}
		}
	}
}

func add(x, y int) {
	tot++
	to[tot] = y
	nxt[tot] = head[x]
	head[x] = tot
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
