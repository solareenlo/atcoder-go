package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := [305][305]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	b := [305][305]int{}
	cnt := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i][j] != 0 {
				if b[j][j] == 0 {
					for k := j; k <= m; k++ {
						b[j][k] = a[i][k]
					}
					cnt++
					break
				} else {
					for k := j; k <= m; k++ {
						a[i][k] ^= b[j][k]
					}
				}
			}
		}
	}
	ans := (powMod(2, n) - powMod(2, n-cnt) + mod) % mod * powMod(2, m-1) % mod
	fmt.Println(ans)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
