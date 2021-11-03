package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, c int
	var s string
	fmt.Scan(&n, &k, &c, &s)
	s = "1" + s

	a := make([]int, 0)
	for i := 1; i <= n; i++ {
		if s[i] == 'o' {
			a = append(a, i)
			i += c
		}
	}
	sort.Ints(a)

	b := make([]int, 0)
	for i := n; i >= 1; i-- {
		if s[i] == 'o' {
			b = append(b, i)
			i -= c
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	for i := 0; i < k; i++ {
		if a[i] == b[k-i-1] {
			fmt.Println(a[i])
		}
	}
}
