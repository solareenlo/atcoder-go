package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	w := make([]float64, n)
	p := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i], &p[i])
	}

	ok, ng, eps := 0.0, 100.0, 1e-8
	for ng-ok > eps {
		m := (ok + ng) / 2
		tmp := make([]float64, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, w[i]*(p[i]-m))
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(tmp)))
		sum := 0.0
		for i := 0; i < k; i++ {
			sum += tmp[i]
		}
		if sum >= 0.0 {
			ok = m
		} else {
			ng = m
		}
	}
	fmt.Println(ok)
}
