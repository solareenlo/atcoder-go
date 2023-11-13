package main

import "fmt"

func main() {
	var p [3][7]int
	for i := 0; i < 3; i++ {
		for j := 1; j <= 6; j++ {
			fmt.Scan(&p[i][j])
		}
	}
	var sum [19]int
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				sum[i+j+k] += p[0][i] * p[1][j] * p[2][k]
			}
		}
	}
	for i := 1; i <= 18; i++ {
		x := float64(sum[i]) / 1000000.0
		fmt.Println(x)
	}
}
