package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c%gcd(a, b) != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
