package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	h := make([]int, n)
	for i := range h {
		fmt.Scan(&h[i])
	}
	sort.Ints(h)

	sum := 0
	for i := 0; i < n-k; i++ {
		sum += h[i]
	}

	fmt.Println(sum)
}
