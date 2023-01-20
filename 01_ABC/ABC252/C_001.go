package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]float64, 15)
	b := make([][]float64, 15)
	for i, _ := range b {
		b[i] = make([]float64, 15)
	}

	for i := 0; i < n; i++ {
		var ss string
		fmt.Scan(&ss)
		for j := 0; j <= 9; j++ {
			k := ss[j] - 48
			a[k] = math.Max(a[k], float64(j)+b[k][j]*10)
			b[k][j]++
		}
	}

	s := 1e9
	for i := 0; i < 10; i++ {
		s = math.Min(s, a[i])
	}
	fmt.Println(s)
}
