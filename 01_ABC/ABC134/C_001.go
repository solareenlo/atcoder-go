package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		b[i] = a[i]
	}
	sort.Ints(b)

	for i := 0; i < n; i++ {
		if a[i] == b[n-1] {
			fmt.Println(b[n-2])
		} else {
			fmt.Println(b[n-1])
		}
	}
}
