package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(2 * (a*b + a*c + b*c))
}
