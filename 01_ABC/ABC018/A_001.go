package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	p := make([]int, 0)
	p = append(p, a)
	p = append(p, b)
	p = append(p, c)
	tmp := make([]int, len(p))
	_ = copy(tmp, p)
	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})

	for i := 0; i < 3; i++ {
		if tmp[i] == p[0] {
			fmt.Println(1)
		}
		if tmp[i] == p[1] {
			fmt.Println(2)
		}
		if tmp[i] == p[2] {
			fmt.Println(3)
		}
	}
}
