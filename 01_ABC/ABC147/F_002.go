package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x, d int
	fmt.Scan(&n, &x, &d)

	if d == 0 {
		if x == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(n + 1)
		}
		return
	}

	if d < 0 {
		x, d = -x, -d
	}

	type pair struct{ a, b int }
	m := make(map[int][]pair)
	for i := 0; i <= n; i++ {
		a := i * (i - 1) / 2
		b := i * (2*n - i - 1) / 2
		c := i * x % d
		m[c] = append(m[c], pair{i*x/d + a, 1}, pair{i*x/d + b + 1, -1})
	}

	res := 0
	for _, e := range m {
		sort.Slice(e, func(i, j int) bool {
			return e[i].a < e[j].a
		})
		sum, pre := 0, 0
		for _, v := range e {
			if sum > 0 {
				res += v.a - pre
			}
			pre = v.a
			sum += v.b
		}
	}
	fmt.Println(res)
}
