package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(3 * (n - gcd(n, m)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
