package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct{ x, y int }

var (
	dx  = []int{0, 1, 0, -1}
	dy  = []int{1, 0, -1, 0}
	mp  = make([]map[node]node, 4)
	vis = map[node]bool{}
)

func opr(c byte) int {
	switch c {
	case 'U':
		return 0
	case 'R':
		return 1
	case 'D':
		return 2
	case 'L':
		return 3
	}
	return -1
}

func move(p node, op int) node {
	if !vis[p] {
		vis[p] = true
		return p
	}
	if _, ok := mp[op][p]; ok {
		return move(mp[op][p], op)
	}
	mp[op][p] = node{0, 0}
	mp[op][p] = move(node{p.x + dx[op], p.y + dy[op]}, op)
	return mp[op][p]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	cur := node{0, 0}
	vis[cur] = true
	for i := range mp {
		mp[i] = make(map[node]node, 1)
	}
	for i := range s {
		cur = move(cur, opr(s[i]))
	}

	fmt.Println(cur.x, cur.y)
}
