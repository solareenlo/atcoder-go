package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const N = 1505
	str := make([]string, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &str[i])
		str[i] = " " + str[i]
	}

	now := 0
	s := [N][N]int{}
	d := [N][N]int{}
	l := [2][N][N]int{}
	r := [2][N][N]int{}
	ans := 0
	for i := n; i > 0; i-- {
		now ^= 1
		for j := n; j > 0; j-- {
			if str[i][j]^'#' != 0 {
				s[i][j] = s[i+1][j] + s[i][j+1] + int(str[i][j]-'0')
				d[i][j] = max(max(d[i+1][j], d[i][j+1]), i)
				l[now][j][i] = j
				r[now][j][i] = j
				for k := i + 1; k <= d[i+1][j]; k++ {
					l[now][j][k] = l[now^1][j][k]
				}
				for k := max(d[i+1][j], i) + 1; k <= d[i][j+1]; k++ {
					l[now][j][k] = l[now][j+1][k]
				}
				for k := i; k <= d[i][j+1]; k++ {
					r[now][j][k] = r[now][j+1][k]
				}
				for k := max(d[i][j+1], i) + 1; k <= d[i+1][j]; k++ {
					r[now][j][k] = r[now^1][j][k]
				}
				k := i + 1
				dw := min(d[i+1][j], d[i][j+1])
				for k <= dw {
					if r[now^1][j][k] >= l[now][j+1][k] {
						x := l[now][j+1][k]
						s[i][j] -= s[k][x]
						k = d[k][x]
					}
					k++
				}
				ans += int(str[i][j]-'0') * (s[i][j] - int(str[i][j]-'0'))
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
