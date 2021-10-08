package main

import (
	"fmt"
	"sort"
)

func main() {
	n := make([]int, 3)
	var k int
	fmt.Scan(&n[0], &n[1], &n[2], &k)
	sort.Ints(n)

	for i := 0; i < k; i++ {
		n[2] *= 2
	}
	fmt.Println(n[0] + n[1] + n[2])
}
