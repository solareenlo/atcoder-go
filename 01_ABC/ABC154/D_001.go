package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	s := make([]float64, n+1)
	s[0] = 0
	for i := 0; i < n; i++ {
		var a float64
		fmt.Scan(&a)
		s[i+1] = s[i] + (1.0+a)/2.0
	}

	res := 0.0
	for i := 0; i < n+1-k; i++ {
		res = max(res, s[i+k]-s[i])
	}
	fmt.Println(res)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
