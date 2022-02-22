package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2002
const P = 1_000_000_007

var (
	n int
)

func upt(o []int, x, v int) {
	for ; x < n+2; x += x & -x {
		o[x] = (o[x] + v) % P
	}
}

func sum(o []int, x int) int {
	res := 0
	for ; x > 0; x -= x & -x {
		res = (res + o[x]) % P
	}
	return res
}

func query(o []int, x, y int) int {
	return (sum(o, y) - sum(o, x) + P) % P
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	a := make([]int, n+2)
	a[n+1] = n + 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	I := make([]int, n+1)
	I[1] = 1
	for i := 2; i <= n; i++ {
		I[i] = (P - P/i) * I[P%i] % P
	}

	dp := [N][N]int{}
	C := make([][]int, N)
	for i := range C {
		C[i] = make([]int, N)
	}
	for j := 1; j < n+2; j++ {
		A := make([]int, 4*n+8)
		B := make([]int, 4*n+8)
		for i := j - 1; i >= 0; i-- {
			if a[i] < a[j] {
				dp[i][j] = (dp[i][j] + query(B, a[i], a[j]) + query(C[i], a[i], a[j])) % P
				k := query(A, a[i], a[j])
				if k != 0 {
					dp[i][j] = (dp[i][j]*I[k] + 1) % P
				}
			}
			upt(C[i], a[j], dp[i][j])
			if i != 0 {
				upt(A, a[i], 1)
				upt(B, a[i], dp[i][j])
			}
		}
	}

	fmt.Println(dp[0][n+1])
}
