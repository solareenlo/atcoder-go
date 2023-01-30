package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5010
	const mod = 998244353

	var n, k int
	fmt.Fscan(in, &n, &k)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	var f [N][N]int
	var c [N]int
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			tmp := 0
			if j != 0 {
				tmp = f[i-1][j-1] * (i - j - c[a[i]] + mod) % mod
			}
			f[i][j] = (f[i-1][j]*(j+c[a[i]]+1)%mod + tmp) % mod
		}
		c[a[i]]++
	}
	fmt.Println(f[n][k])
}
