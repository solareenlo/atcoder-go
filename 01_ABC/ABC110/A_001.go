package main

import (
	"fmt"
	"sort"
)

func main() {
	n := make([]int, 3)
	fmt.Scan(&n[0], &n[1], &n[2])
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println(n[0]*10 + n[1] + n[2])
}
