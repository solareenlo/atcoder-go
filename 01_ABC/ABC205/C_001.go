package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if int(c)%2 == 0 {
		a = math.Abs(a)
		b = math.Abs(b)
	}

	if a < b {
		fmt.Println("<")
	} else if a > b {
		fmt.Println(">")
	} else {
		fmt.Println("=")
	}
}
