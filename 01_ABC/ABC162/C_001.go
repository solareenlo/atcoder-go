package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	sum := 0
	for k := 1; k < K+1; k++ {
		for j := 1; j < K+1; j++ {
			for i := 1; i < K+1; i++ {
				sum += gcd(gcd(i, j), k)
			}
		}
	}

	fmt.Println(sum)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
