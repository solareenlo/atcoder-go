package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	e := [12][12]int{}
	for x, y, i := 0, 0, 0; i < m; i++ {
		fmt.Scan(&x, &y)
		e[x-1][y-1] = 1
	}

	res := 1
	var i, j, t int
	for bit := 0; bit < (1 << n); bit++ {
		for t, i = 1, 0; i < n; i++ {
			for j = i + 1; j < n; j++ {
				t &= f((bit>>i)&1) | f((bit>>j)&1) | e[i][j]
			}
		}
		for i, j = 0, 0; i < n; i++ {
			j += (bit >> i) & 1
		}
		if res < j*t {
			res = j * t
		}
	}
	fmt.Println(res)
}

func f(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
