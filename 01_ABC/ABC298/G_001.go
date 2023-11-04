package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 7
const N2 = 6*6 + 2
const N4 = 5e2 + 2

var vis [N][N][N][N]bool
var pre [N][N]int
var cnt, t int
var val []int
var f [N][N][N][N][N2][N4]int

func dfs(u, d, l, r int) {
	if vis[u][d][l][r] {
		return
	}

	sum := pre[d][r] - pre[d][l-1] - pre[u-1][r] + pre[u-1][l-1]
	for i := 1; i <= cnt; i++ {
		if sum >= val[i] {
			f[u][d][l][r][0][i] = sum
		}
	}

	for i := 1; i <= t && i <= (r-l+1)*(d-u+1)-1; i++ {
		for j := u; j < d; j++ {
			dfs(u, j, l, r)
			dfs(j+1, d, l, r)
			for k := 0; k < i; k++ {
				for p := 1; p <= cnt; p++ {
					f[u][d][l][r][i][p] = min(f[u][d][l][r][i][p], max(f[u][j][l][r][k][p], f[j+1][d][l][r][i-k-1][p]))
				}
			}
		}
		for j := l; j < r; j++ {
			dfs(u, d, l, j)
			dfs(u, d, j+1, r)
			for k := 0; k < i; k++ {
				for p := 1; p <= cnt; p++ {
					f[u][d][l][r][i][p] = min(f[u][d][l][r][i][p], max(f[u][d][l][j][k][p], f[u][d][j+1][r][i-k-1][p]))
				}
			}
		}
	}
	vis[u][d][l][r] = true
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	val = make([]int, N*N*N*N)

	var n, m int
	fmt.Fscan(in, &n, &m, &t)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &pre[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre[i][j] = pre[i][j] + pre[i][j-1]
		}
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			pre[i][j] = pre[i][j] + pre[i-1][j]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := i; k <= n; k++ {
				for l := j; l <= m; l++ {
					cnt++
					val[cnt] = pre[k][l] - pre[i-1][l] - pre[k][j-1] + pre[i-1][j-1]
				}
			}
		}
	}

	sort.Ints(val[1 : cnt+1])
	const INF = 4557430888798830399
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for ii := 0; ii < N; ii++ {
					for jj := 0; jj < N2; jj++ {
						for kk := 0; kk < N4; kk++ {
							f[i][j][k][ii][jj][kk] = INF
						}
					}
				}
			}
		}
	}
	dfs(1, n, 1, m)

	ans := int(2e18) + 7
	for i := 1; i <= cnt; i++ {
		ans = min(ans, f[1][n][1][m][t][i]-val[i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
