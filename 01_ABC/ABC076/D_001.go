package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	t, v := make([]int, n), make([]int, n)
	for i := range t {
		fmt.Scan(&t[i])
		t[i] *= 2
	}
	for i := range v {
		fmt.Scan(&v[i])
		v[i] *= 2
	}

	j, to := 0, 0
	z := make([]int, 200*200+1)
	for i := 0; i < n; i++ {
		to += t[i]
		z[j] = min(z[j], v[i])
		for j++; j <= to; j++ {
			z[j] = min(z[j-1]+1, v[i])
		}
		j--
	}

	mv, sum := 0, 0
	for i := j - 1; i >= 0; i-- {
		mv = min(z[i], mv+1)
		sum += mv
	}
	fmt.Println(float64(sum) / 4.0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
