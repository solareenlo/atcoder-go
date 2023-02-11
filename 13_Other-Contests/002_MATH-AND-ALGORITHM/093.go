package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	G := gcd(A, B)
	if A/G > 1000000000000000000/B {
		fmt.Println("Large")
	} else {
		fmt.Println(A / G * B)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
