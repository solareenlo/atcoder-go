package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, s, k int
		fmt.Scan(&n, &s, &k)

		g, x, _ := extGcd(k, n)

		if s%g != 0 {
			fmt.Println(-1)
		} else {
			n /= g
			s /= g
			k /= g
			fmt.Println(((x * -s % n) + n) % n)
		}
	}
}

func extGcd(x, y int) (int, int, int) {
	if y == 0 {
		return x, 1, 0
	}
	g, a, b := extGcd(y, x%y)
	return g, b, a - x/y*b
}
