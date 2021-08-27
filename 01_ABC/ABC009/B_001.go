package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 1001)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		a[tmp] = tmp
	}
	sort.Ints(a)

	fmt.Println(a[len(a)-2])
}
