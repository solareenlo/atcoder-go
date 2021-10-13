package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	sort.Ints(p)
	p[n-1] /= 2

	sum := 0
	for i := range p {
		sum += p[i]
	}
	fmt.Println(sum)
}
