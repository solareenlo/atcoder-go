package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, H, M float64
	fmt.Scan(&A, &B, &H, &M)

	theta := math.Min(360-math.Abs(30*H-5.5*M), math.Abs(30*H-5.5*M)) * math.Pi / 180.0
	fmt.Println(math.Sqrt(A*A + B*B - 2.0*A*B*math.Cos(theta)))
}
