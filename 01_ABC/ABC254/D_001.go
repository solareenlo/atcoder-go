package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	for i := 1; i*i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				a += n / (i * i) * 2
			}
		}
	}

	fmt.Println(a + n)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
