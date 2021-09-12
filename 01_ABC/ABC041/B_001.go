package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	mod := int(1e9 + 7)
	fmt.Println(a * b % mod * c % mod)
}
