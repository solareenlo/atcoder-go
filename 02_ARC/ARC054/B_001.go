package main

import (
	"fmt"
	"math"
)

func main() {
	var p float64
	fmt.Scan(&p)

	a := 1.5 / math.Log(2)
	if p < a {
		fmt.Println(p)
	} else {
		fmt.Println(-1.5*(math.Log2(a)-math.Log2(p)) + a)
	}
}
