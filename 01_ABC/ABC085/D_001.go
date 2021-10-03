package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h int
	fmt.Scan(&n, &h)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	index := sort.SearchInts(b, a[n-1])

	res := 0
	for i := n - 1; i >= index; i-- {
		h -= b[i]
		res++
		if h <= 0 {
			break
		}
	}

	if h > 0 {
		div := h / a[n-1]
		rem := h % a[n-1]
		if rem != 0 {
			res += div + 1
		} else {
			res += div
		}
	}
	fmt.Println(res)
}
