package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println((b-a)*((c-b-1)/(b-a)+1) + b)
}
