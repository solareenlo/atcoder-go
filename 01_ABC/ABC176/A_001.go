package main

import "fmt"

func main() {
	var n, x, t int
	fmt.Scan(&n, &x, &t)

	div := n / x
	rem := n % x

	if rem != 0 {
		fmt.Println(t * (div + 1))
	} else {
		fmt.Println(t * div)
	}
}
