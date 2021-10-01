package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	c := make([][]int, 10)
	for i := range c {
		c[i] = make([]int, 10)
		for j := range c[i] {
			fmt.Scan(&c[i][j])
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if c[i][j] > c[i][k]+c[k][j] {
					c[i][j] = c[i][k] + c[k][j]
				}
			}
		}
	}

	res, a := 0, 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Scan(&a)
			if a >= 0 {
				res += c[a][1]
			}
		}
	}
	fmt.Println(res)
}
