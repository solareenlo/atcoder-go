package main

import "fmt"

func main() {
	var q, a, b int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		fmt.Scan(&a, &b)
		if a > b {
			a, b = b, a
		}
		mul := a * b
		l, r := a, b+1
		for r-l > 1 {
			mid := (r + l) / 2
			c := (mid + a) / 2
			if mid == a {
				c--
			}
			if c*(a+mid-c) < mul {
				l = mid
			} else {
				r = mid
			}
		}
		fmt.Println(l + a - 2)
	}
}
