package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(abs(a[0]-a[1]) + abs(a[1]-a[2]))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
