package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 5005

	var f [N][N]int

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+3)
	for i := 1; i <= n+1; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a[1 : n+1])
	sort.Ints(b[1 : n+2])
	if n == 1 {
		if a[1] < b[1] && a[2] < b[2] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}
	f[1][1] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := -1; k <= 1; k++ {
				if j+k < 0 || i+j+k > n {
					continue
				}
				if k >= 0 && a[i+1] >= b[i+j+1] {
					continue
				}
				now := j + MOD
				if k == 0 {
					now *= 2
				}
				f[i+1][j+k] = (f[i+1][j+k] + now*f[i][j]) % MOD
			}
		}
	}
	fmt.Println(f[n][1])
}
