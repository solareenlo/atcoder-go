package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := [][]int{[]int{1, 1, 1}, []int{1, 0, 0}, []int{0, 1, 0}}
	a = matPowMod(a, n-1)
	ans := (a[2][0]*2 + a[2][1]*1 + a[2][2]*1) % mod
	fmt.Println(ans)
}

const mod = 1_000_000_007

func matPowMod(a [][]int, n int) [][]int {
	ans := [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}
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
	ans := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ans[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ans[i][j] = (ans[i][j] + a[i][k]*b[k][j]%mod) % mod
			}
		}
	}
	return ans
}
