package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	type node struct{ x, id int }
	a := make([]node, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x)
		a[i].id = i
	}

	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	p := make([]int, n+1)
	pid := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		pid[p[i]] = i
	}

	for i := 1; i <= n; i++ {
		if i != p[i] && a[i].x <= b[p[i]] {
			fmt.Fprintln(out, -1)
			return
		}
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x < tmp[j].x
	})

	m := 0
	aa := make([]int, n+1)
	ab := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if p[a[i].id] != a[i].id {
			now := a[i].id
			m++
			aa[m] = now
			ab[m] = pid[now]
			p[pid[now]] = p[now]
			pid[p[now]] = pid[now]
		}
	}

	fmt.Fprintln(out, m)
	for i := 1; i <= m; i++ {
		fmt.Fprintln(out, aa[i], ab[i])
	}
}
