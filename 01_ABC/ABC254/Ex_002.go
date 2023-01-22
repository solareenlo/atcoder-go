package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var a, b [N]int
var t [N << 5][2]int
var s [N << 5]int
var ans int
var cnt int

func insert(v, f int) {
	p, i := 1, 0
	for (1 << i) <= v {
		i++
	}
	for i >= 0 {
		w := (v >> i) & 1
		if t[p][w] == 0 {
			cnt++
			t[p][w] = cnt
		}
		p = t[p][w]
		i--
	}
	s[p] += f
}

func dfs(x int) {
	if x == 0 {
		return
	}
	ls := t[x][0]
	rs := t[x][1]
	dfs(ls)
	ans += abs(s[ls])
	s[x] += s[ls]
	dfs(rs)
	ans += abs(s[rs])
	s[x] += s[rs]
	if s[rs] < 0 {
		fmt.Println(-1)
		os.Exit(0)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	cnt = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		insert(a[i], 1)
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		insert(b[i], -1)
	}
	dfs(1)
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
