package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	f, s := 0, 0
	for i := 0; i < n; i += 2 {
		if a[i] == a[i+1] {
			s = f
			f = a[i]
		} else {
			i--
		}
	}
	fmt.Println(s * f)
}
