package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	type pair struct{ w, v int }
	p := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].w, &p[i].v)
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].v > p[j].v
	})

	x := make([]int, m)
	for i := range x {
		fmt.Scan(&x[i])
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		l--
		r--
		s := make([]int, 0)
		for j := 0; j < m; j++ {
			if j < l || r < j {
				s = append(s, x[j])
			}
		}
		sort.Ints(s)
		res := 0
		for _, x := range p {
			pos := lowerBound(s, x.w)
			if len(s) != 0 && pos < len(s) {
				res += x.v
				s = erase(s, pos)
			}
		}
		fmt.Println(res)
	}
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
