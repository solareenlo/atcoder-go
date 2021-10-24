package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	base := make([]int, 0)

	all := 0
	for i := range a {
		all ^= a[i]
	}
	for i := range a {
		a[i] &= ^all
	}

	for i := range a {
		for j := range base {
			a[i] = min(a[i], a[i]^base[j])
		}
		if a[i] != 0 {
			base = append(base, a[i])
			sort.Sort(sort.Reverse(sort.IntSlice(base)))
		}
	}

	up := 0
	for i := range base {
		up = max(up, up^base[i])
	}
	fmt.Println(all + 2*up)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
