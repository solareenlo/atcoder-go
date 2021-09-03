package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	if a/b == c/d {
		fmt.Println("DRAW")
	} else if a/b > c/d {
		fmt.Println("AOKI")
	} else {
		fmt.Println("TAKAHASHI")
	}
}
