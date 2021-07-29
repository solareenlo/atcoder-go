package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	if sum < k {
		fmt.Println(-1)
		return
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	res := 0
	for i := range a {
		k -= a[i]
		res++
		if k < 1 {
			break
		}
	}
	fmt.Println(res)
}
