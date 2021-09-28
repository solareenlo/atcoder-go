package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var res, t int
	fmt.Scan(&res)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&t)
		g := gcd(res, t)
		res /= g
		res *= t
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
