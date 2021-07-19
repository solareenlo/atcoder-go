package main

import "fmt"

func main() {
	var n, a, x, y int
	fmt.Scan(&n, &a, &x, &y)
	if n > a {
		fmt.Println(a*x + (n-a)*y)
	} else {
		fmt.Println(n * x)
	}
}
