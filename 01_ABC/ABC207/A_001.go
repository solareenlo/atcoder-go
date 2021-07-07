package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a[:])
	fmt.Println(a[1] + a[2])
}
