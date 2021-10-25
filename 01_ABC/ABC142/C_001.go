package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	type pair struct{ a, index int }
	s := make([]pair, 0)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		s = append(s, pair{a, i + 1})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].a < s[j].a
	})

	for i := 0; i < n; i++ {
		fmt.Print(s[i].index)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
