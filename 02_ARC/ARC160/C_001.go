package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 500010
	const MOD = 998244353

	var a [N]int
	var f [N][]int

	var n int
	fmt.Fscan(in, &n)
	mx := 0
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x]++
		mx = max(mx, x)
	}
	f[1] = append(f[1], 1)
	for i := 1; i <= mx+log2(n); i++ {
		siz := len(f[i])
		resize(&f[i+1], (siz+a[i])/2+1)
		for j := siz - 1; j >= 0; j-- {
			if j != siz-1 {
				f[i][j] = (f[i][j] + f[i][j+1]) % MOD
			}
			f[i+1][(j+a[i])/2] = (f[i+1][(j+a[i])/2] + f[i][j]) % MOD
		}
	}
	fmt.Println(f[mx+log2(n<<1)][0])
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
