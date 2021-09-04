package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]float64, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	res := -1e9
	for i := 0; i < n; i++ {
		maxA, maxT := -1e9, -1e9
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			l := int(math.Min(float64(i), float64(j)))
			r := int(math.Max(float64(i), float64(j)))
			sumA, sumT := 0.0, 0.0
			for k := l; k <= r; k++ {
				if (k-l)%2 == 1 {
					sumA += a[k]
				} else {
					sumT += a[k]
				}
			}
			if sumA > maxA {
				maxA, maxT = sumA, sumT
			}
		}
		res = math.Max(res, maxT)
	}
	fmt.Println(res)
}
