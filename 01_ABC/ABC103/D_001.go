package main

import (
	"fmt"
	"sort"
)

type pair struct {
	a, b int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	ba := make([]pair, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&ba[i].b, &ba[i].a)
	}
	sort.Slice(ba, func(i, j int) bool {
		return ba[i].a < ba[j].a
	})

	cnt, c := 0, 0
	for i := range ba {
		if c < ba[i].b {
			cnt++
			c = ba[i].a - 1
		}
	}
	fmt.Println(cnt)
}
