package main

import "fmt"

func main() {
	var N, K, M int
	fmt.Scan(&N, &K, &M)
	if N%(M/gcd(K, M)) != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
