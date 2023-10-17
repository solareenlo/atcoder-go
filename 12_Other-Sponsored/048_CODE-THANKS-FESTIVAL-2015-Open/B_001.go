package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C, D, E int
	fmt.Scan(&A, &B, &C, &D, &E)
	ans := make(map[int]bool)
	if A == E || B == E {
		ans[C] = true
		ans[D] = true
	}
	if C == E || D == E {
		ans[A] = true
		ans[B] = true
	}
	fmt.Println(len(ans))
	keys := make([]int, 0, len(ans))
	for k := range ans {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, i := range keys {
		fmt.Println(i)
	}
}
