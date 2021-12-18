package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	mp := map[float64]bool{}
	t := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if y[i] != y[j] {
				mp[float64(x[i]-x[j])/float64(y[i]-y[j])] = true
			} else {
				t = 1
			}
		}
	}
	fmt.Println(len(mp)*2 + t*2)
}
