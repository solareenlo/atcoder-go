package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var b = make([]int, n)
	copy(b, a)
	sort.Ints(b)

	m := k / n
	for i := 0; i < n; i++ {
		if a[i] < b[k-n*m] {
			fmt.Println(m + 1)
		} else {
			fmt.Println(m)
		}
	}
}
