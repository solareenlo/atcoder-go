package main

import (
	"fmt"
	"math"
)

func main() {
	var y1, y2, y3 float64
	fmt.Println("? 1")
	fmt.Scan(&y1)
	fmt.Println("? 2")
	fmt.Scan(&y2)
	fmt.Println("? 3")
	fmt.Scan(&y3)
	g := -y3 + 2*y2 - y1
	v := (y2 - y1 + g*1.5) * math.Sqrt(2)
	h := y1 + g/2 - v/math.Sqrt(2)
	fmt.Println("!", h, v, g)
}
