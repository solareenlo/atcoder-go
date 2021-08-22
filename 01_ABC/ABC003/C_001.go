package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	r := make([]int, n)
	for i := range r {
		fmt.Scan(&r[i])
	}
	sort.Ints(r)
	r = r[n-k:]

	res := 0.0
	for i := range r {
		res = (res + float64(r[i])) / 2
	}
	fmt.Println(res)
}
