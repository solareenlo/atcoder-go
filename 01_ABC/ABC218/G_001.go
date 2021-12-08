package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

var (
	n int
	a = [N]int{}
	b = make([]int, N)
	c = [N]int{}
	g = make([][]int, N)
)

func mdy(i, w int) {
	for ; i <= n; i += i & -i {
		c[i] += w
	}
}

func qry(x int) int {
	p := 0
	for i := 16; i >= 0; i-- {
		if p+(1<<i) < n && c[p|(1<<i)] < x {
			p |= 1 << i
			x -= c[p]
		}
	}
	return b[p+1]
}

func dfs(u, p, d int) int {
	mdy(a[u], 1)
	w := 1 << 60
	if d&1 != 0 {
		w = 0
	}
	wv := 0
	if len(g[u]) == 1 && u != 1 {
		w = (qry(d/2) + qry(d/2+1)) / 2
		if d&1 != 0 {
			w = qry((d + 1) / 2)
		}
	}
	for _, v := range g[u] {
		if v != p {
			wv = dfs(v, u, d+1)
			if d&1 > 0 {
				w = max(w, wv)
			} else {
				w = min(w, wv)
			}
		}
	}
	mdy(a[u], -1)
	return w
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
	}
	sort.Ints(b[1 : n+1])

	for i := 1; i < n+1; i++ {
		pos := lowerBound(b[1:n+1], a[i])
		a[i] = pos + 1
	}

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	fmt.Println(dfs(1, 0, 1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
