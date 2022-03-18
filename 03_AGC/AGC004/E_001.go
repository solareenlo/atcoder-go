package main

import "fmt"

var (
	sr = [110][110]int{}
	sc = [110][110]int{}
)

func qr(x, l, r int) int { return sr[x][r] - sr[x][l-1] }
func qc(y, l, r int) int { return sc[r][y] - sc[l-1][y] }

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var x, y int
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		s = " " + s
		for j := 1; j <= m; j++ {
			w := 0
			if s[j] == 'o' {
				w = 1
			}
			sr[i][j] = sr[i][j-1] + w
			sc[i][j] = sc[i-1][j] + w
			if s[j] == 'E' {
				x = i
				y = j
			}
		}
	}

	f := [110][110][110]int{}
	for u, rd := 1, n-x+1; u <= x; u, rd = u+1, rd+1 {
		for d, ru := n, n-x+1; d >= x; d, ru = d-1, ru-1 {
			for l, rr := 1, m-y+1; l <= y; l, rr = l+1, rr+1 {
				for r, rl := m, m-y+1; r >= y; r, rl = r-1, rl-1 {
					if u-1 >= ru {
						f[d][l][r] = max(f[d][l][r], f[d][l][r]+qr(u-1, max(l, rl), min(r, rr)))
					}
					if d+1 <= rd {
						f[d][l][r] = max(f[d][l][r], f[d+1][l][r]+qr(d+1, max(l, rl), min(r, rr)))
					}
					if l-1 >= rl {
						f[d][l][r] = max(f[d][l][r], f[d][l-1][r]+qc(l-1, max(u, ru), min(d, rd)))
					}
					if r+1 <= rr {
						f[d][l][r] = max(f[d][l][r], f[d][l][r+1]+qc(r+1, max(u, ru), min(d, rd)))
					}
				}
			}
		}
	}
	fmt.Println(f[x][y][y])
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
