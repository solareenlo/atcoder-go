package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	first, second int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	adj := make([][]pair, n)
	for i := 1; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		adj[a] = append(adj[a], pair{b, c})
		adj[b] = append(adj[b], pair{a, c})
	}

	dis := make([]int, n)
	bfs := func(s int) int {
		for i := range dis {
			dis[i] = -1
		}
		dis[s] = 0
		q := []int{s}

		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			for _, p := range adj[x] {
				y, w := p.first, p.second
				if dis[y] == -1 {
					dis[y] = dis[x] + w
					q = append(q, y)
				}
			}
		}

		maxIndex := 0
		for i := 1; i < n; i++ {
			if dis[maxIndex] < dis[i] {
				maxIndex = i
			}
		}

		return maxIndex
	}

	s := bfs(0)
	t := bfs(s)
	ds := make([]int, n)
	copy(ds, dis)
	bfs(t)
	dt := dis

	if s > t {
		s, t = t, s
		ds, dt = dt, ds
	}

	for i := 0; i < n; i++ {
		var ans int
		if ds[i] >= dt[i] {
			ans = s + 1
		} else {
			ans = t + 1
		}
		fmt.Fprintln(out, ans)
	}
}
