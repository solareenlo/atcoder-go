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
	fmt.Println(a[1])
}
