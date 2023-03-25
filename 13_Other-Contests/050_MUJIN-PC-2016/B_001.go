package main

import (
	"fmt"
	"sort"
)

const N = 1e6 + 5

func main() {
	var a [4]int
	fmt.Scan(&a[1], &a[2], &a[3])
	sort.Ints(a[:4])
	x := float64(a[1]+a[2]+a[3]) * float64(a[1]+a[2]+a[3])
	r := float64(a[3] - a[2] - a[1])
	if r > 0 {
		x -= r * r
	}
	x *= 3.1415926535897932
	fmt.Printf("%.10f", x)
}
