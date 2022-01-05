package main

import (
	"fmt"
	"sort"
)

func main() {
	b := map[string]int{}
	for i := 0; i < 10; i++ {
		var tmp string
		fmt.Scan(&tmp)
		b[tmp] = i
	}

	type P struct {
		s string
		x int
	}

	var n int
	fmt.Scan(&n)
	a := make([]P, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		a[i] = P{tmp, 0}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i].s); j++ {
			a[i].x *= 10
			a[i].x += b[string(a[i].s[j])]
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].s < a[j].s)
	})

	for i := 0; i < n; i++ {
		fmt.Println(a[i].s)
	}
}
