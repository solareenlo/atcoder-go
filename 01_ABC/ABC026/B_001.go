package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	r := make([]int, n)
	for i := range r {
		fmt.Scan(&r[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	res := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			res += r[i] * r[i]
		} else {
			res -= r[i] * r[i]
		}
	}
	fmt.Println(float64(res) * math.Pi)
}
