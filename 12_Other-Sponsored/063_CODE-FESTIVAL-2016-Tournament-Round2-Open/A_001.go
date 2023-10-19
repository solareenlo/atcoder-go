package main

import "fmt"

func main() {
	var x int
	var p float64
	fmt.Scan(&x, &p)
	fmt.Println(50.0 * float64(x+x%2) / p)
}
