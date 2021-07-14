package main

import (
	"fmt"
	"sort"
)

type pair struct {
	t int
	s string
}

func main() {
	var n int
	fmt.Scan(&n)
	m := make([]pair, n)
	for i := range m {
		fmt.Scan(&m[i].s)
		fmt.Scan(&m[i].t)
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].t > m[j].t
	})

	fmt.Println(m[1].s)
}
