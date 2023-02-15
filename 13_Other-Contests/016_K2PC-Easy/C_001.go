package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	i := 1
	a, b, c, d := 0, 0, 0, 0
	for i > 0 {
		i++
		for j := 1; j < i; j++ {
			d++
			if d == n || d == m {
				a += i - j
				b += j
				c++
				if n == m {
					a += i - j
					b += j
					c++
				}
			}
			if a == i-j && b == j && c == 2 {
				fmt.Println(d)
				return
			}
		}
	}
}
