package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]float64, 2001)
	y := make([]float64, 2001)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	rad := make([]float64, 4001)
	epsilon := 1e-10
	a, b := 0, 0
	for i := 0; i < n; i++ {
		l := 0
		for j := 0; j < n; j++ {
			if i != j {
				rad[l] = math.Atan2(x[i]-x[j], y[i]-y[j])
				l++
			}
		}
		sort.Float64s(rad[:l])
		for j := 0; j < n-1; j++ {
			rad[l] = rad[l-n+1] + math.Pi*2
			l++
		}
		r := 0
		for j := 0; j < n-1; j++ {
			for rad[r]-rad[j] < math.Pi/2-epsilon {
				r++
			}
			if rad[r]-rad[j] < math.Pi/2+epsilon {
				b++
				r++
			}
			a += n - 1 - r
		}
	}
	fmt.Println(n*(n-1)*(n-2)/6-a-b, b, a)
}
