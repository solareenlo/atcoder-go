package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(1)
		return
	}

	res := [6]int{2, 4, 8, 16, 32, 64}
	index := sort.Search(len(res), func(i int) bool { return n < res[i] })
	fmt.Println(res[index-1])
}
