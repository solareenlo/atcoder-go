package main

import "fmt"

func main() {
	var n, C int
	fmt.Scan(&n, &C)
	D := [30][30]int{}
	for y := 0; y < C; y++ {
		for x := 0; x < C; x++ {
			fmt.Scan(&D[y][x])
		}
	}
	c := [501][501]int{}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Scan(&c[y][x])
			c[y][x]--
		}
	}

	pre := [3][30]int{}
	for col := 0; col < C; col++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				pre[(y+x)%3][col] += D[c[y][x]][col]
			}
		}
	}

	res := 1 << 60
	for i := 0; i < C; i++ {
		for j := 0; j < C; j++ {
			for k := 0; k < C; k++ {
				if i == j || j == k || k == i {
					continue
				}
				res = min(res, pre[0][i]+pre[1][j]+pre[2][k])
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
