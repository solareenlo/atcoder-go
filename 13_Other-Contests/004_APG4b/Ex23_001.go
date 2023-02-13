package main

import (
	"fmt"
	"sort"
)

type pair struct {
	f, s int
}

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		m[a]++
	}

	v := make([]pair, 0)
	for K, V := range m {
		v = append(v, pair{K, V})
	}
	sort.Slice(v, func(i, j int) bool {
		if v[i].s != v[j].s {
			return v[i].s < v[j].s
		}
		return v[i].f < v[j].f
	})

	fmt.Println(v[len(v)-1].f, v[len(v)-1].s)
}
