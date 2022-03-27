package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	fmt.Println((k%2*2 - 1) * (b - a))
}
