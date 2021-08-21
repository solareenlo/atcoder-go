package main

import "fmt"

func main() {
	var xa, ya, xb, yb, xc, yc float64
	fmt.Scan(&xa, &ya, &xb, &yb, &xc, &yc)

	res := abs((xa*yb + xb*yc + xc*ya - ya*xb - yb*xc - yc*xa) / 2)
	fmt.Printf("%.5f\n", res)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
