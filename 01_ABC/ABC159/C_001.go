package main

import "fmt"

func main() {
	var l float64
	fmt.Scan(&l)

	n := l / 3.0
	fmt.Println(n * n * n)
}
