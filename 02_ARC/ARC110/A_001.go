package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 1
	for i := 2; i <= n; i++ {
		res = res * i / gcd(res, i)
	}
	fmt.Println(res + 1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
