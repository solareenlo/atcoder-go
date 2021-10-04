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
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	sumA, sumB := 0, 0
	for i := 0; i < len(a); i++ {
		if i%2 != 0 {
			sumB += a[i]
		} else {
			sumA += a[i]
		}
	}
	fmt.Println(sumA - sumB)
}
