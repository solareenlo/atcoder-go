package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	var N, M int
	fmt.Scan(&R, &N, &M)

	var a [100005]float64
	for i := 1; i < N; i++ {
		dist := math.Abs(float64(i)-float64(N)/2) / (float64(N) / 2) * R
		a[i] = math.Sqrt(1.0*R*R-dist*dist) * 2
	}
	ans := 0.0
	for i := N + M - 1; i >= 1; i-- {
		mx := 0.0
		if i < N {
			mx = a[i]
		}
		if i-M >= 1 && i-M <= N {
			mx = math.Max(mx, a[i-M])
		}
		ans += mx
	}
	fmt.Println(ans)
}
