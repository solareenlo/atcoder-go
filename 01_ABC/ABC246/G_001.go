package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	m int
	a = make([]int, 202020)
	b = make([]int, 202020)
	q = make([][]int, 202020)
)

func dfs(x, fa int) int {
	re := 0
	for i := 1; i < len(q[x])+1; i++ {
		if q[x][i-1] != fa {
			re += dfs(q[x][i-1], x)
		}
	}
	tmp := 0
	if re != 0 {
		tmp = 1
	}
	tmp2 := 0
	if a[x] >= b[m] {
		tmp2 = 1
	}
	return re - tmp + tmp2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 2; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 2; i < n+1; i++ {
		b[i] = a[i]
	}
	for i := 2; i < n+1; i++ {
		var s, l int
		fmt.Fscan(in, &s, &l)
		q[s] = append(q[s], l)
		q[l] = append(q[l], s)
	}
	tmp := b[1 : n+1]
	sort.Ints(tmp)
	s := 1
	l := n
	for s < l {
		m = (s + l + 1) >> 1
		if dfs(1, 0) != 0 {
			s = m
		} else {
			l = m - 1
		}
	}
	fmt.Println(b[s])
}
