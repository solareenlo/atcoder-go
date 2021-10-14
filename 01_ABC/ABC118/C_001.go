package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var tmp int
	fmt.Scan(&tmp)
	g := gcd(tmp, tmp)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&tmp)
		g = gcd(g, tmp)
	}
	fmt.Println(g)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
