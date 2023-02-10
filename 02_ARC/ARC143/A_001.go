package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	sort.Ints(a)

	if a[2] > (a[1] + a[0]) {
		fmt.Println(-1)
	} else {
		fmt.Println(a[2])
	}
}
