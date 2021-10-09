package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n+1; i++ {
		a[i] += a[i-1]
	}

	x, y := 0, 2
	res := 1 << 60
	for i := 2; i < n-1; i++ {
		for a[x+1] < a[i]-a[x+1] {
			x++
		}
		for a[y+1]-a[i] < a[n]-a[y+1] {
			y++
		}
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				p := a[x+j]
				q := a[i] - a[x+j]
				r := a[y+k] - a[i]
				s := a[n] - a[y+k]
				maxi := max(p, q, r, s)
				mini := min(p, q, r, s)
				res = min(res, abs(maxi-mini))
			}
		}
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
