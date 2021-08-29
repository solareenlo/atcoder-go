package main

import "fmt"

func main() {
	var n, h, a, b, c, d, e int
	fmt.Scan(&n, &h, &a, &b, &c, &d, &e)

	res := int(1e16)
	for i := 0; i <= n; i++ {
		l, r := -1, n-i
		for r-l > 1 {
			m := (l + r) / 2
			x := h + b*i + d*m - (n-i-m)*e
			if x > 0 {
				r = m
			} else {
				l = m
			}
		}
		res = min(res, a*i+c*r)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
