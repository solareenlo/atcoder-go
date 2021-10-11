package main

import "fmt"

type p struct {
	x, y int
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := [501][501]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	l, r := make([]p, 0), make([]p, 0)
	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			if a[i][j]%2 != 0 {
				a[i][j+1]++
				l = append(l, p{i + 1, j + 1})
				r = append(r, p{i + 1, j + 1 + 1})
			}
		}
		if i != h-1 && a[i][w-1]%2 != 0 {
			a[i+1][w-1]++
			l = append(l, p{i + 1, w - 1 + 1})
			r = append(r, p{i + 1 + 1, w - 1 + 1})
		}
	}

	n := len(l)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println(l[i].x, l[i].y, r[i].x, r[i].y)
	}
}
