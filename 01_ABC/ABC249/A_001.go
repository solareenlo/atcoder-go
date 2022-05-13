package main

import "fmt"

func main() {
	var a, b, c, d, e, f, x int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &x)

	m := (x/(a+c)*a + min(x%(a+c), a)) * b
	n := (x/(d+f)*d + min(x%(d+f), d)) * e

	if m > n {
		fmt.Println("Takahashi")
	} else if m < n {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
