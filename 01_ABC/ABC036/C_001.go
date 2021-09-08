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
	}
	copy(b, a)
	sort.Ints(b)

	m := make(map[int]int, n)
	j := 0
	for i := 0; i < n; i++ {
		if _, ok := m[b[i]]; !ok {
			m[b[i]] = j
			j++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(m[a[i]])
	}
}
