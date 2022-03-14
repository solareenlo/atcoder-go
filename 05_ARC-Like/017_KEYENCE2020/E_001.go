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

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 100007
	p := make([]int, n+1)
	d := make([]int, n+1)
	type pair struct{ x, y int }
	tmp := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i])
		tmp[i] = pair{d[i], i}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x < tmp[j].x
	})
	for i := 1; i < n+1; i++ {
		p[i] = tmp[i].y
	}

	f := make([]int, N)
	w := make([]int, N*2)
	e := make([]int, N)
	id := make([]int, N)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if d[u] == d[v] {
			if f[u] == 0 && f[v] == 0 {
				f[u] = 1
			}
			if f[u] == 0 {
				f[u] = -f[v]
			}
			if f[v] == 0 {
				f[v] = -f[u]
			}
			if f[u] == -f[v] {
				w[i] = d[u]
			}
		} else {
			if d[u] < d[v] {
				u, v = v, u
			}
			e[u] = v
			id[u] = i
		}
	}

	for i := 1; i <= n; i++ {
		x := p[i]
		if f[x] == 0 {
			if e[x] == 0 {
				fmt.Fprintln(out, -1)
				return
			}
			f[x] = f[e[x]]
			w[id[x]] = d[x] - d[e[x]]
		}
	}

	for i := 1; i <= n; i++ {
		if f[i] > 0 {
			fmt.Fprint(out, "B")
		} else {
			fmt.Fprint(out, "W")
		}
	}
	fmt.Fprintln(out)
	for i := 1; i <= m; i++ {
		if w[i] != 0 {
			fmt.Fprintln(out, w[i])
		} else {
			fmt.Fprintln(out, 1000000000)
		}
	}
}
