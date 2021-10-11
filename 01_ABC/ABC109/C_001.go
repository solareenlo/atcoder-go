package main

import "fmt"

func main() {
	var n, X int
	fmt.Scan(&n, &X)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}

	if n == 1 {
		fmt.Println(abs(X - x[0]))
		return
	}

	g := gcd(abs(X-x[0]), abs(X-x[0]))
	for i := 0; i < n; i++ {
		g = gcd(g, abs(X-x[i]))
	}
	fmt.Println(abs(g))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
