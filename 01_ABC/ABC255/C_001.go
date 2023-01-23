package main

import "fmt"

func main() {
	var x, a, d, n int
	fmt.Scan(&x, &a, &d, &n)

	b := a + d*n - d
	s := 0
	if d == 0 {
		s = abs(x - a)
	} else {
		if b < a {
			a, b = b, a
			d = -d
		}
		if x < a {
			s = a - x
		} else if x > b {
			s = x - b
		} else {
			x = (x - a) % d
			s = min(x, d-x)
		}
	}

	fmt.Println(s)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
