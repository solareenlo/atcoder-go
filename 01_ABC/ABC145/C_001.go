package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	town := make([]int, 0)
	for i := 0; i < n; i++ {
		town = append(town, i)
	}

	tmp, sum := n, 1
	for tmp != 1 {
		sum *= tmp
		tmp--
	}

	dist := 0.0
	for i := 0; i < sum; i++ {
		for j := 0; j < n-1; j++ {
			dist += math.Hypot(x[town[j]]-x[town[j+1]], y[town[j]]-y[town[j+1]])
		}
		nextPermutation(sort.IntSlice(town))
	}

	fmt.Printf("%.10f\n", dist/float64(sum))
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
