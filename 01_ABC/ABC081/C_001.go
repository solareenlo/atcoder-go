package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n+1)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		a[tmp]++
	}
	a = a[:n]
	sort.Ints(a)

	res := 0
	for i := n - 1; i >= n-k; i-- {
		res += a[i]
	}
	fmt.Println(n - res)
}
