package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := 0.0
	a := make([]float64, 2*n+1)
	for i := 1; i <= 2*n; i++ {
		fmt.Scan(&a[i])
	}
	tmp := a[1 : 2*n+1]
	sort.Float64s(tmp)
	for i := 1; i <= 2*n; i++ {
		ans += a[i] * (float64(i) - 1.0) * 1.0 / (2*float64(n) - 1)
	}
	fmt.Println(ans)
}
