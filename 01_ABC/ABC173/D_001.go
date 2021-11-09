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
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	res := a[0]
	cnt := 1
	i := 1
	for {
		if cnt == n-1 {
			break
		}
		res += a[i]
		cnt++
		if cnt == n-1 {
			break
		}
		res += a[i]
		cnt++
		i++
	}
	fmt.Println(res)
}
