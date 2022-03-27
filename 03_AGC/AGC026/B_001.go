package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		if a < b || d < b || (a-c-1)/gcd(b, d) > (a-b)/gcd(b, d) {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
