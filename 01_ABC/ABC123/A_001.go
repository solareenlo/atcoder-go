package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 5)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	var k int
	fmt.Scan(&k)

	if a[4]-a[0] <= k {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
