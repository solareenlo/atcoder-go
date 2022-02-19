package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Scan(&k, &n)

	const N = 5050
	const M = 5000
	const p = 998244353
	C := [N][N]int{}
	C[0][0] = 1
	for i := 1; i <= M; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < M; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % p
		}
	}

	for i := 2; i <= 2*k; i++ {
		x := i/2 - i + k + 1
		if i <= k+1 {
			x = i / 2
		}
		ans := 0
		for j := 0; j <= x; j++ {
			tmp := 1
			if j&1 != 0 {
				tmp = p - 1
			}
			ans = (ans + tmp*C[x][j]%p*C[n-2*j+k-1][k-1]) % p
		}
		fmt.Fprintln(out, ans)
	}
}
