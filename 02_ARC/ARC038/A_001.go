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

	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += a[i]
		}
	}

	fmt.Println(sum)
}
