package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))

	res := 0
	for i := 0; i < k; i++ {
		res += l[i]
	}
	fmt.Println(res)
}
