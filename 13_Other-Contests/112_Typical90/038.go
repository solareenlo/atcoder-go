package main

import "fmt"

func main() {
	const INF = int(1e18)

	var a, b int
	fmt.Scan(&a, &b)

	x := gcd(a, b)
	ans := a / x
	if INF/b < ans {
		fmt.Println("Large")
	} else {
		fmt.Println(ans * b)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
