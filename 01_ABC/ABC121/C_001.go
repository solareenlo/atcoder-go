package main

import (
	"fmt"
	"sort"
)

type pair struct{ a, b int }

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	drink := make([]pair, n)
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		drink[i] = pair{a, b}
	}

	sort.Slice(drink, func(i, j int) bool {
		return drink[i].a < drink[j].a
	})

	res, i := 0, 0
	for m > 0 {
		if drink[i].b <= m {
			res += drink[i].a * drink[i].b
			m -= drink[i].b
		} else {
			res += m * drink[i].a
			break
		}
		i++
	}
	fmt.Println(res)
}
