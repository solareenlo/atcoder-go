package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := [][]int{[]int{1, 1}, []int{1, 0}}
	a = matPowMod(a, n-1)
	ans := (a[1][0]*1 + a[1][1]*1) % mod
	fmt.Println(ans)
}

const mod = 1_000_000_000

func matPowMod(a [][]int, n int) [][]int {
	ans := [][]int{[]int{1, 0}, []int{0, 1}}
	for n != 0 {
		if n%2 == 1 {
			ans = matMulMod(ans, a)
		}
		a = matMulMod(a, a)
		n /= 2
	}
	return ans
}

func matMulMod(a, b [][]int) [][]int {
	ans := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ans[i] = make([]int, 2)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				ans[i][j] = (ans[i][j] + a[i][k]*b[k][j]%mod) % mod
			}
		}
	}
	return ans
}
