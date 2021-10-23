package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	for i := range v {
		fmt.Scan(&v[i])
	}
	sort.Ints(v)

	res := float64(v[0])
	for i := 0; i < n-1; i++ {
		res = (res + float64(v[i+1])) / 2.0
	}
	fmt.Println(res)
}
