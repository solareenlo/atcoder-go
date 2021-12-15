package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]float64, n)
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	t := 0.0
	for i := 0; i < n; i++ {
		t += a[i] / b[i]
	}
	t /= 2.0

	res := 0.0
	for i := 0; i < n; i++ {
		res += min(a[i], t*b[i])
		t -= min(a[i]/b[i], t)
	}

	fmt.Println(res)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
