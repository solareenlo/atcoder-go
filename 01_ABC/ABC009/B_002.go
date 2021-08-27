package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	set := make(map[int]struct{})
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		set[tmp] = struct{}{}
	}

	keys := make([]int, 0)
	for k := range set {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println(keys[len(keys)-2])
}
