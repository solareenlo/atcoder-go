package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var t []int
	for i := 1; i < n+1; i++ {
		t = append(t, i)
	}

	var ask func([]int) int
	ask = func(t []int) int {
		if len(t) == 0 {
			return 0
		}
		fmt.Fprintf(out, "? %d", len(t))
		out.Flush()
		for _, x := range t {
			fmt.Fprintf(out, " %d", x)
			out.Flush()
		}
		fmt.Fprintln(out)
		out.Flush()
		var k int
		fmt.Fscan(in, &k)
		return k
	}
	m := ask(t)

	var deg [255]int
	var vis [255]bool
	q := make([]int, 0)
	for i := 1; i < n+1; i++ {
		t = make([]int, 0)
		for j := 1; j < n+1; j++ {
			if i^j != 0 {
				t = append(t, j)
			}
		}
		deg[i] = m - ask(t)
		if deg[i] != 0 {
			vis[i] = true
		}
		if deg[i] == 1 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if !vis[x] {
			continue
		}
		vis[x] = false
		t = make([]int, 0)
		for i := 1; i < n+1; i++ {
			if vis[i] {
				t = append(t, i)
			}
		}
		var dfs func(int, []int) int
		dfs = func(x int, t []int) int {
			if len(t) == 1 {
				return t[0]
			}
			tl := make([]int, 0)
			tr := make([]int, 0)
			for i := 0; i < len(t)/2; i++ {
				tl = append(tl, t[i])
			}
			for i := len(t) / 2; i < len(t); i++ {
				tr = append(tr, t[i])
			}
			ptl := make([]int, len(tl))
			copy(ptl, tl)
			ptl = append(ptl, x)
			if ask(tl) == ask(ptl) {
				return dfs(x, tr)
			}
			return dfs(x, tl)
		}
		y := dfs(x, t)
		deg[y]--
		if deg[y] == 0 {
			vis[y] = false
		}
		if deg[y] == 1 {
			q = append(q, y)
		}
	}
	flg := false
	for i := 1; i < n+1; i++ {
		if !flg && !(vis[i] && deg[i] > 1) {
			flg = false
		} else {
			flg = true
		}
	}
	if flg {
		fmt.Fprintln(out, "! Yes")
		out.Flush()
	} else {
		fmt.Fprintln(out, "! No")
		out.Flush()
	}
}
