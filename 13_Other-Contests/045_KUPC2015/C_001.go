package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		d := make([][]int, n)
		for i := 0; i < n; i++ {
			d[i] = make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Scan(&d[i][j])
				if d[i][j] == -1 {
					d[i][j] = 1 << 29
				}
			}
		}
		res := 1
		for i := 0; i < n; i++ {
			if d[i][i] != 0 {
				res = 0
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					if d[i][j]+d[j][k] < d[i][k] {
						res = 0
					}
				}
			}
		}
		if res == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
