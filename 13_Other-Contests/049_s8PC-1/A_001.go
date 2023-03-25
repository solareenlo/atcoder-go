package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%.12f\n", 1080.0/float64(n-1))
}
