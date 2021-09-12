package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		m[a[i]] = i + 1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for i := 0; i < n; i++ {
		fmt.Println(m[a[i]])
	}
}
