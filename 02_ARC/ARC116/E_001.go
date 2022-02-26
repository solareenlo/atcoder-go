package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 200200

type Edge struct{ nxt, to int }

var (
	edge = [maxN << 1]Edge{}
	cnt  int
	mid  int
	N    int
	head = make([]int, maxN)
	n    int
	m    int
)

func add_edge(x, y int) {
	N++
	edge[N] = Edge{head[x], y}
	head[x] = N
	N++
	edge[N] = Edge{head[y], x}
	head[y] = N
}

func dfs(x, y int) int {
	up := -314514
	dow := 0
	for i := head[x]; i > 0; i = edge[i].nxt {
		des := edge[i].to
		if des == y {
			continue
		}
		tt := dfs(des, x)
		if tt < 0 {
			dow = max(dow, -tt)
		} else {
			up = max(up, tt)
		}
	}
	up--
	if up >= dow {
		return up
	} else if dow == mid {
		cnt++
		return mid
	} else {
		return -(dow + 1)
	}
}

func chk() bool {
	cnt = 0
	if dfs(1, 0) < 0 {
		cnt++
	}
	return cnt <= m
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i < n; i++ {
		var t, k int
		fmt.Fscan(in, &t, &k)
		add_edge(t, k)
	}

	L := 0
	R := n - 1
	ans := 0
	for L <= R {
		mid = (L + R) / 2
		if chk() {
			ans = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
