package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	x := 1.0
	for i := 10 - 1; i >= 0; i-- {
		y := a * x * x * x * x
		x -= (x*(y+b) + c) / (5*y + b)
	}
	fmt.Println(x)
}
