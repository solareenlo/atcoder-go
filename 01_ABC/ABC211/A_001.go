package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println((a-b)/3 + b)
}
