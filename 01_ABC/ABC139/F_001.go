package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	s, t := [100]float64{}, [100]float64{}
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &t[i])
	}

	maxi := 0.0
	for i := 0; i < 360; i++ {
		x := 0.0
		y := 0.0
		for j := 0; j < n; j++ {
			rad := float64(i) * math.Pi / 180.0
			if math.Cos(rad)*s[j]+math.Sin(rad)*t[j] > 0.0 {
				x += s[j]
				y += t[j]
			}
		}
		maxi = max(maxi, x*x+y*y)
	}

	fmt.Println(math.Sqrt(maxi))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
