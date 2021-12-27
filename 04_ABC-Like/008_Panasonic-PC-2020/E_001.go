package main

import "fmt"

func ok(x, y byte) bool {
	return x == '?' || y == '?' || x == y
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	n := 100000
	ab := make([]bool, n)
	ac := make([]bool, n)
	bc := make([]bool, n)
	A := len(a)
	B := len(b)
	C := len(c)
	for i := 0; i < A; i++ {
		for j := 0; j < B; j++ {
			if !ok(a[i], b[j]) {
				ab[i-j+50000] = true
			}
		}
	}
	for i := 0; i < A; i++ {
		for j := 0; j < C; j++ {
			if !ok(a[i], c[j]) {
				ac[i-j+50000] = true
			}
		}
	}
	for i := 0; i < B; i++ {
		for j := 0; j < C; j++ {
			if !ok(b[i], c[j]) {
				bc[i-j+50000] = true
			}
		}
	}
	res := 6000
	for i := -4000; i <= 4000; i++ {
		for j := -4000; j <= 4000; j++ {
			if !(ab[i+50000] || ac[j+50000] || bc[j-i+50000]) {
				res = min(res, max(A, max(B+i, C+j))-min(0, min(i, j)))
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
