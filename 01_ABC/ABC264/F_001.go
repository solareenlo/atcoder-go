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
	var h, l [2009]int
	var a [2009][2009]byte
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
		a[i][0] = '0'
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &l[i])
		a[0][i] = '0'
	}
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := range s {
			a[i][j+1] = s[j]
		}
	}

	const MAXI = 4_557_430_888_798_830_399
	var f [2009][2009][2][2]int
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				for l := range f[i][j][k] {
					f[i][j][k][l] = MAXI
				}
			}
		}
	}
	f[1][1][0][0] = 0
	f[1][1][0][1] = l[1]
	f[1][1][1][0] = h[1]
	f[1][1][1][1] = h[1] + l[1]

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for p := 0; p <= 1; p++ {
				for q := 0; q <= 1; q++ {
					f[i][j][p][q] = min(
						min(f[i][j][p][q], f[i-1][j][p^int(a[i][j])^int(a[i-1][j])][q]+p*h[i]), f[i][j-1][p][q^int(a[i][j])^int(a[i][j-1])]+q*l[j])
				}
			}
		}
	}

	ans := MAXI
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			ans = min(f[n][m][i][j], ans)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
