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
	sort.Sort(sort.IntSlice(a))
	for i := range a {
		fmt.Print(a[i], " ")
	}
}
