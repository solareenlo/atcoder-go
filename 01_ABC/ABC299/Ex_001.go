package main

import "fmt"

func main() {
	const mod = 998244353
	const inv6 = (mod + 1) / 6

	var m [7][7]int

	var x int
	fmt.Scan(&x)
	for i := 0; i < 5; i++ {
		m[i+1][i] = 1
	}
	for i := 0; i < 6; i++ {
		m[i][5] = inv6
	}
	m[6][5] = 1
	m[6][6] = 1
	ans := [7]int{915136369, 514943221, 192518427, 860371429, 261480227, 0, 1}
	for i := 0; i < 30; i++ {
		if (x & (1 << i)) != 0 {
			var val [7]int
			for j := 0; j < 7; j++ {
				for k := 0; k < 7; k++ {
					val[k] = (ans[j]*m[j][k] + val[k]) % mod
				}
			}
			for j := 0; j < 7; j++ {
				ans[j] = val[j]
			}
		}
		var p [7][7]int
		for j := 0; j < 7; j++ {
			for k := 0; k < 7; k++ {
				for l := 0; l < 7; l++ {
					p[j][l] = (p[j][l] + m[j][k]*m[k][l]) % mod
				}
			}
		}
		for j := 0; j < 7; j++ {
			for k := 0; k < 7; k++ {
				m[j][k] = p[j][k]
			}
		}
	}
	fmt.Println(ans[5])
}
