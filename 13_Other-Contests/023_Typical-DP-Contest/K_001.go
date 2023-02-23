package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	p := make([]pair, n)
	for i := 0; i < n; i++ {
		var x, r int
		fmt.Fscan(in, &x, &r)
		p[i] = pair{x + r, x - r}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})

	d := make([]int, n)
	for i := range d {
		d[i] = 1e9
	}
	for i := 0; i < n; i++ {
		idx := lowerBound(d, -p[i].y)
		d[idx] = -p[i].y
	}
	fmt.Println(lowerBound(d, 1e9))
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
