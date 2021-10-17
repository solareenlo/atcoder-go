package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	l := make([]int, n)
	for i := 0; i < n-1; i++ {
		l[i+1] = gcd(l[i], a[i])
	}

	r := make([]int, n)
	for i := n - 1; i > 0; i-- {
		r[i-1] = gcd(r[i], a[i])
	}

	res := 1
	for i := 0; i < n; i++ {
		res = max(res, gcd(l[i], r[i]))
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
