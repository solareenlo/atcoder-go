package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type edge struct{ x, y int }

var (
	p   []edge
	e   = make([][]edge, 100005)
	res []int
	vis = [100005]bool{}
	a   = [100005]int{}
)

func dfs(x int) int {
	vis[x] = true
	ret := a[x]
	for _, v := range e[x] {
		y := v.x
		if vis[y] {
			continue
		}
		t := dfs(y)
		if t != 0 {
			res = append(res, v.y)
		}
		ret ^= t
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	p = make([]edge, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	a[0] = p[0].y
	for i := 1; i < n; i++ {
		a[i] = p[i-1].y ^ p[i].y
	}
	a[n] = p[n-1].y

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l = lowerBound(p, edge{l, 0})
		r = upperBound(p, edge{r, 1})
		e[l] = append(e[l], edge{r, i + 1})
		e[r] = append(e[r], edge{l, i + 1})
	}

	for i := 0; i < n+1; i++ {
		if !vis[i] {
			if dfs(i) != 0 {
				fmt.Fprintln(out, -1)
				return
			}
		}
	}

	fmt.Fprintln(out, len(res))
	sort.Ints(res)
	for i := range res {
		fmt.Fprint(out, res[i])
		if i != len(res)-1 {
			fmt.Fprint(out, " ")
		}
	}
}

func lowerBound(a []edge, x edge) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

func upperBound(a []edge, x edge) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x > x.x
	})
	return idx
}
