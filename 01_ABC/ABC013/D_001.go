package main

import "fmt"

func main() {
	var n, m, d, t int
	fmt.Scan(&n, &m, &d)

	a := make([]int, n)
	c := make([]int, n)
	for i := range a {
		a[i] = i
		c[i] = i
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&t)
		a[t-1], a[t] = a[t], a[t-1]
	}

	b := make([]int, n)
	for ; d > 0; d /= 2 {
		if d%2 != 0 {
			for i := 0; i < n; i++ {
				c[i] = a[c[i]]
			}
		}
		for i := 0; i < n; i++ {
			b[i] = a[i]
		}
		for i := 0; i < n; i++ {
			a[i] = b[a[i]]
		}
	}

	res := make([]int, n)
	for i := range res {
		res[c[i]] = i
	}
	for i := range res {
		fmt.Println(res[i] + 1)
	}
}
