package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	c := [101][101]int{}
	for i := 0; i < h*w; i++ {
		fmt.Scan(&c[i/w][i%w])
	}

	x := [101][101]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if (i+j)&1 != 0 {
				x[i+1][j+1] = x[i+1][j] + x[i][j+1] - x[i][j] - c[i][j]
			} else {
				x[i+1][j+1] = x[i+1][j] + x[i][j+1] - x[i][j] + c[i][j]
			}
		}
	}

	res := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := i + 1; k <= h; k++ {
				for l := j + 1; l <= w; l++ {
					if x[k][l]-x[i][l]-x[k][j]+x[i][j] == 0 {
						res = max(res, (k-i)*(l-j))
					}
				}
			}
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
