package main

import "fmt"

func main() {
	var n, k, x, y int
	fmt.Scan(&n, &k, &x, &y)

	if n > k {
		fmt.Println(k*x + (n-k)*y)
	} else {
		fmt.Println(n * x)
	}
}
