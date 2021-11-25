package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	children = make([][]int, 200002)
	list     = make([][]int, 200002)
	in       = make([]int, 200002)
	out      = make([]int, 200002)
	depth    = make([]int, 200002)
	timer    int
)

func dfs(u int) {
	in[u] = timer
	timer++
	list[depth[u]] = append(list[depth[u]], in[u])
	for _, v := range children[u] {
		depth[v] = depth[u] + 1
		dfs(v)
	}
	out[u] = timer
	timer++
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(stdin, &N)

	for i := 1; i < N; i++ {
		var p int
		fmt.Fscan(stdin, &p)
		children[p-1] = append(children[p-1], (i))
	}

	dfs(0)

	var Q int
	fmt.Fscan(stdin, &Q)
	for i := 0; i < Q; i++ {
		var u, d int
		fmt.Fscan(stdin, &u, &d)
		u -= 1
		v := list[d]
		fmt.Println(lowerBound(v, out[u]) - lowerBound(v, in[u]))
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
