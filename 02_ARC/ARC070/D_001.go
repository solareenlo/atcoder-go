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
	sort.Ints(a)

	sum, res := 0, n
	for i := n - 1; i >= 0; i-- {
		if sum+a[i] < k {
			sum += a[i]
		} else {
			res = i
		}
	}
	fmt.Println(res)
}
