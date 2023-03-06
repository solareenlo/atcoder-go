package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	cnt, sum := 0, 0
	for sum < n {
		sum += a[cnt%3]
		cnt++
	}
	fmt.Println(cnt)
}
