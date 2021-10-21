package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	c := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
		c[i] = p[i]
	}
	sort.Ints(p)

	cnt := 0
	for i := 0; i < n; i++ {
		if p[i] != c[i] {
			cnt++
		}
	}

	if cnt == 2 || cnt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
