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
	v := make([]pair, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		v = append(v, pair{a, b})
	}
	sortPair(v)
	w := make([]int, 0)
	w = append(w, 0)
	for i := 0; i < n-1; i++ {
		w = append(w, v[i].x+v[i+1].y-v[i].y)
	}
	var m int
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		var t int
		fmt.Fscan(in, &t)
		u := upperBound(w, t) - 1
		fmt.Fprintln(out, v[u].y+max(0, t-v[u].x))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
