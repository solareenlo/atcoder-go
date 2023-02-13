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

	ab := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ab[i].f, &ab[i].s)
	}
	sort.Slice(ab, func(i, j int) bool {
		if ab[i].s != ab[j].s {
			return ab[i].s < ab[j].s
		}
		return ab[i].f < ab[j].f
	})

	for i := 0; i < n; i++ {
		fmt.Println(ab[i].f, ab[i].s)
	}
}
