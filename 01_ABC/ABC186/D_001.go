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
	sort.Ints(a)

	res := 0
	for i := 0; i < n; i++ {
		res += (i+1)*a[i] - (n-i)*a[i]
	}

	fmt.Println(res)
}
