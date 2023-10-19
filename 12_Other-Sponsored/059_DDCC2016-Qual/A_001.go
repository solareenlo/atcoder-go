package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%.16f\n", c*1.0*b/a)
}
