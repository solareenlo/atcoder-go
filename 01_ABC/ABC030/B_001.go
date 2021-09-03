package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	nDeg := float64((90-n*30)%360) - float64(m)/2.0
	mDeg := float64((90 - m*6) % 360)
	diff := math.Abs(mDeg - nDeg)
	fmt.Println(math.Min(diff, 360.0-diff))
}
