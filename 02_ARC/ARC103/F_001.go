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

	a := make([]int, n+1)
	rv := map[int]int{}
	const N = 100005
	sz := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		rv[a[i]] = i
		sz[i] = 1
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	fa := make([]int, N)
	for i := n; i > 1; i-- {
		u := rv[a[i]]
		f := rv[a[i]+sz[u]*2-n]
		fa[u] = f
		sz[f] += sz[u]
		if f == 0 || fa[f] != 0 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	dep := make([]int, N)
	dst := 0
	for i := 2; i <= n; i++ {
		u := rv[a[i]]
		dep[u] = dep[fa[u]] + 1
		dst += dep[u]
	}

	if dst != a[1] {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 2; i <= n; i++ {
		u := rv[a[i]]
		fmt.Fprintln(out, u, fa[u])
	}
}
