package main

import (
	"fmt"
	"sort"
)

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	if n%2 == 1 {
		fmt.Println(-1)
	} else {
		v := make([]float64, n)
		for i := 0; i < n; i++ {
			var x, y int
			fmt.Scan(&x, &y)
			v[i] = float64(x) / float64(y)
		}
		sort.Float64s(v)
		flag := false
		for i := 1; i <= w; i++ {
			for j := 1; j <= h; j++ {
				if i == w || j == h {
					if v[n/2-1] < float64(i)/float64(j) && v[n/2] > float64(i)/float64(j) {
						fmt.Printf("(%d,%d)\n", i, j)
						flag = true
					}
				}
			}
		}
		if !flag {
			fmt.Println(-1)
		}
	}
}
