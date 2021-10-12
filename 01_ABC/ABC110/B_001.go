package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, X, Y int
	fmt.Scan(&n, &m, &X, &Y)

	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(x)))

	y := make([]int, m)
	for i := range y {
		fmt.Scan(&y[i])
	}
	sort.Ints(y)

	if x[0] >= y[0] || x[0] >= Y || y[0] <= X {
		fmt.Println("War")
	} else {
		fmt.Println("No War")
	}
}
