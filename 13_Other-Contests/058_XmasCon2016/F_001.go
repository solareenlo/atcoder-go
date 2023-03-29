package main

import "fmt"

func main() {
	var t, s, n, a, b, k int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		s = 0
		fmt.Scan(&n, &a, &b, &k)
		for p := 0; p <= n; p++ {
			if p <= a && n-p <= b {
				if k == 1 {
					if p == 0 || n-p == 0 {
						s++
					}
				} else {
					if (n & p) == p {
						s++
					}
				}
			}
		}
		fmt.Println(s % 2)
	}
}
