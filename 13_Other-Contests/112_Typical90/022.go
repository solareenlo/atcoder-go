package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	d := gcd(a, gcd(b, c))
	fmt.Println((a/d - 1) + (b/d - 1) + (c/d - 1))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
