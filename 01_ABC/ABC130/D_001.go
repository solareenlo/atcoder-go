package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] += s[i] + a[i]
	}

	sum := 0
	for i := 0; i < n; i++ {
		pos := sort.SearchInts(s, s[i]+k)
		sum += n - pos + 1
	}
	fmt.Println(sum)
}
