package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 333

	var b, c, pre, stk, l, r [N]int
	var a [N][N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[j] = 0
			c[j] = 310
		}
		for j := i; j <= n; j++ {
			for k := 1; k <= m; k++ {
				b[k] += a[j][k]
				c[k] = min(c[k], a[j][k])
				pre[k] = pre[k-1] + b[k]
			}
			cnt := 1
			stk[cnt] = 0
			for k := 1; k <= m; k++ {
				for cnt != 0 && c[stk[cnt]] >= c[k] {
					cnt--
				}
				l[k] = stk[cnt] + 1
				cnt++
				stk[cnt] = k
			}
			cnt = 0
			cnt++
			stk[cnt] = m + 1
			for k := m; k >= 1; k-- {
				for cnt != 0 && c[stk[cnt]] >= c[k] {
					cnt--
				}
				r[k] = stk[cnt] - 1
				cnt++
				stk[cnt] = k
				ans = max(ans, c[k]*(pre[r[k]]-pre[l[k]-1]))
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
