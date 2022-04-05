package main

import "fmt"

func main() {
	var a, b, c, x float64
	fmt.Scan(&a, &b, &c, &x)

	p := 0.0
	if x <= a {
		p = 1
	}
	if x > a && x <= b {
		p = c / (b - a)
	}
	fmt.Println(p)
}
