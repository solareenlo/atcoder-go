package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 200200

	var a, b [N]int
	var f [N][10]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	f[1][0] = 1
	f[1][1] = 1
	for i := 2; i <= n; i++ {
		if a[i] != a[i-1] {
			f[i][0] = (f[i][0] + f[i-1][0]) % MOD
		}
		if a[i] != b[i-1] {
			f[i][0] = (f[i][0] + f[i-1][1]) % MOD
		}
		if b[i] != a[i-1] {
			f[i][1] = (f[i][1] + f[i-1][0]) % MOD
		}
		if b[i] != b[i-1] {
			f[i][1] = (f[i][1] + f[i-1][1]) % MOD
		}
	}
	fmt.Println((f[n][0] + f[n][1]) % MOD)
}
