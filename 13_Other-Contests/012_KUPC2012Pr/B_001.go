package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	for i := range a {
		fmt.Print(a[i])
	}
	fmt.Println()
}
