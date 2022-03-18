package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const mod = 1_000_000_007

var res int

func fastmi(x, n int) int {
	if n == 0 {
		return 1
	}
	if n&1 != 0 {
		res = fastmi(x, n>>1)
		return res * res % mod * x % mod
	}
	res = fastmi(x, n>>1)
	return res * res % mod
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	k %= (mod - 1)

	a := make([]string, n+1)
	a[0] = strings.Repeat(" ", m+2)
	var A, B, C, B1, C1, B2, C2 int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] = " " + a[i] + " "
		for j := 1; j <= m; j++ {
			if a[i][j] == '#' {
				A++
			}
			if a[i-1][j] == '#' && a[i][j] == '#' {
				B2++
			}
			if a[i][j-1] == '#' && a[i][j] == '#' {
				B1++
			}
			if i == n && a[1][j] == '#' && a[i][j] == '#' {
				C2++
			}
		}
		if a[i][1] == '#' && a[i][m] == '#' {
			C1++
		}
	}
	if C1 != 0 && C2 != 0 {
		fmt.Println(1)
		return
	}
	if C1 == 0 && C2 == 0 {
		fmt.Println(fastmi(A, k-1))
		return
	}
	if C1 != 0 && C2 == 0 {
		B = B1
		C = C1
	}
	if C1 == 0 && C2 != 0 {
		B = B2
		C = C2
	}
	fmt.Println(((fastmi(A, k-1)-fastmi(C-A, mod-2)*(fastmi(C, k-1)-fastmi(A, k-1))%mod*B)%mod + mod) % mod)
}
