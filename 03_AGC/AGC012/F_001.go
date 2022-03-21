package main

import (
	"fmt"
	"sort"
)

const P = 1_000_000_007

func add(x, y int) int {
	x += y
	if x >= P {
		x -= P
	}
	return x
}

func main() {
	const N = 109

	var n int
	fmt.Scan(&n)
	m := 2 * n

	a := make([]int, N)
	for i := 1; i < m; i++ {
		fmt.Scan(&a[i])
	}
	tmp := a[1:m]
	sort.Ints(tmp)

	f := [N][N]int{}
	f[0][0] = 1
	for i := n - 1; i > 0; i-- {
		g := [N][N]int{}
		u := 0
		if a[i] != a[i+1] {
			u = 1
		}
		v := 0
		if a[m-i] != a[m-i-1] {
			v = 1
		}
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if f[j][k] != 0 {
					g[j+u][k+v] = add(g[j+u][k+v], f[j][k])
					for l := j + u - 1; l >= 0; l-- {
						g[l][k+v+1] = add(g[l][k+v+1], f[j][k])
					}
					for l := k + v - 1; l >= 0; l-- {
						g[j+u+1][l] = add(g[j+u+1][l], f[j][k])
					}
				}
			}
		}
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				f[j][k] = g[j][k]
			}
		}
	}

	ans := 0
	for j := 0; j < m; j++ {
		for k := 0; k < m; k++ {
			ans = add(ans, f[j][k])
		}
	}
	fmt.Println(ans)
}
