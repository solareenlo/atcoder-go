package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	A := math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
	B := math.Sqrt((a-e)*(a-e) + (b-f)*(b-f))
	C := math.Sqrt((e-c)*(e-c) + (f-d)*(f-d))
	s := (A + B + C) / 2
	S := math.Sqrt(s * (s - A) * (s - B) * (s - C))
	fmt.Println(S * 2 / (A + B + C + 4*S/math.Max(math.Max(A, B), C)))
}
