package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	type pair struct {
		x, y int
	}

	fmt.Fscan(in, &n)
	if n == 1 {
		fmt.Fprintln(out, "!")
		out.Flush()
		return
	}

	root := -1
	for i := 0; i < n; i++ {
		v := make([]int, 1)
		v[0] = i
		if query(v, 1<<(n-1)) == n-1 {
			root = i
			break
		}
	}
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		if i == root {
			continue
		}
		v := make([]int, 0)
		v = append(v, root)
		v = append(v, i)
		x := 3 << (n - 2)
		sz[i] = query(v, x)
	}
	sz[root] = n
	ord := make([]int, n)
	for i := 0; i < n; i++ {
		ord[i] = i
	}
	sort.Slice(ord, func(x, y int) bool {
		return sz[ord[x]] > sz[ord[y]]
	})

	ans := make([]pair, 0)
	used := make([]bool, n)
	used[root] = true
	for _, i := range ord {
		if i == root {
			continue
		}
		for j := 0; j < n; j++ {
			if used[j] && sz[j] > sz[i] {
				v := make([]int, 0)
				v = append(v, j)
				v = append(v, i)
				x := 3 << (n - 2)
				if query(v, x) == sz[i]*(n-sz[i]) {
					ans = append(ans, pair{i, j})
					sz[j] -= sz[i]
					used[i] = true
					break
				}
			}
		}
	}
	fmt.Fprintln(out, "!")
	out.Flush()
	for _, p := range ans {
		fmt.Fprintln(out, p.x+1, p.y+1)
		out.Flush()
	}
}

func query(v []int, x int) int {
	q := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < len(v); i++ {
		q[v[i]] = n - i - 1
		used[v[i]] = true
	}
	idx := 0
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		q[i] = idx
		idx++
	}
	fmt.Fprint(out, "? ")
	out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", q[i]+1)
		out.Flush()
	}
	fmt.Fprintln(out, x)
	out.Flush()
	var ret int
	fmt.Fscan(in, &ret)
	return ret
}
