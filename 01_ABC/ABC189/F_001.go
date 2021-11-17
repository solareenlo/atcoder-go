package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	a := make([]int, 200002)
	for i := 0; i < k; i++ {
		var x int
		fmt.Scan(&x)
		a[x] = 1
	}

	sumA, sumB := 0.0, 0.0
	dpA := make([]float64, 200002)
	dpB := make([]float64, 200002)
	for i := n - 1; i >= 0; i-- {
		if a[i] != 0 {
			dpB[i] = 1
		} else {
			dpA[i] = sumA/float64(m) + 1.0
			dpB[i] = sumB / float64(m)
		}
		sumA += dpA[i] - dpA[i+m]
		sumB += dpB[i] - dpB[i+m]
	}

	if k == 0 {
		fmt.Println(dpA[0])
	} else if dpB[0] < 1.0-1e-8 {
		fmt.Println(dpA[0] / (1.0 - dpB[0]))
	} else {
		fmt.Println(-1)
	}
}
