package main

import (
	"fmt"
	"math"
)

const (
	PI = math.Pi
)

var SIN [200000]float64

func area(i, j int) float64 {
	return 2 * SIN[i] * SIN[j] * SIN[j-i]
}

func main() {
	var n int
	var k int64
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		SIN[i] = math.Sin(PI * float64(i) / float64(n))
	}

	lo, hi := 0.0, 1.3
	for j := 0; j < 35; j++ {
		mi := (lo + hi) / 2
		cnt := int64(0)
		for i := 1; i < n-1; i++ {
			lo2, hi2 := i, (n+i)/2+1
			for hi2-lo2 > 1 {
				mi2 := (lo2 + hi2) / 2
				if area(i, mi2) < mi {
					lo2 = mi2
				} else {
					hi2 = mi2
				}
			}
			l, r := lo2, n-(lo2-i)
			tmp := 0
			if l == r {
				tmp = 1
			}
			cnt += int64((l - i) + (n - r) - tmp)
		}
		cnt = cnt * int64(n) / 3
		if cnt < k {
			lo = mi
		} else {
			hi = mi
		}
	}
	fmt.Printf("%.15f\n", lo)
}
