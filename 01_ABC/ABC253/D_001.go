package main

import "fmt"

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)
	fmt.Println(N*(N+1)/2 - (N/A)*(N/A+1)/2*A - (N/B)*(N/B+1)/2*B + (N/lcm(A, B))*(N/lcm(A, B)+1)/2*lcm(A, B))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
