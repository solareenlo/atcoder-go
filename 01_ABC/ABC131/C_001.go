package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	l := lcm(c, d)
	divL := (b / l) - ((a - 1) / l)
	divC := (b / c) - ((a - 1) / c)
	divD := (b / d) - ((a - 1) / d)

	fmt.Println((b - a + 1) - (divC + divD - divL))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
