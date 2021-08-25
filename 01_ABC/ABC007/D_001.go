package main

import "fmt"

func f(n int64) int64 {
	m, b := int64(0), int64(1)
	for i := n; i > 0; i /= 10 {
		d := i % 10
		if d == 4 || d == 9 {
			m = 0
		}
		if d > 4 {
			m += (d - 1) * b
		} else {
			m += d * b
		}
		b *= 8
	}
	return n - m
}

func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	fmt.Println(f(b+1) - f(a))
}
