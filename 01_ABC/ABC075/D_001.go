package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	res := 1 << 62
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					l, r, u, d := x[a], x[b], y[c], y[d]
					if r < l || u < d {
						continue
					}
					cnt := 0
					for i := 0; i < n; i++ {
						if l <= x[i] && x[i] <= r && d <= y[i] && y[i] <= u {
							cnt++
						}
					}
					if cnt >= k {
						res = min(res, (r-l)*(u-d))
					}
				}
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
