package main

import "fmt"

func main() {
	var R, C, k, n int
	fmt.Scan(&R, &C, &k, &n)

	X := make([]int, 100001)
	Y := make([]int, 100001)
	x := make([]int, 100001)
	y := make([]int, 100001)
	r := make([]int, 100001)
	c := make([]int, 100001)
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i], &Y[i])
		X[i]--
		Y[i]--
		r[X[i]]++
		c[Y[i]]++
	}
	for i := 0; i < R; i++ {
		x[r[i]]++
	}
	for i := 0; i < C; i++ {
		y[c[i]]++
	}
	res := 0
	for i := 0; i <= k; i++ {
		res += x[i] * y[k-i]
	}
	for i := 0; i < n; i++ {
		t := r[X[i]] + c[Y[i]]
		if t == k {
			res--
		} else if t == k+1 {
			res++
		}
	}
	fmt.Println(res)
}
