package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100015

var (
	n  int
	s  string
	t  string
	aa = [N]int{}
	bb = [N]int{}
	dp = [N][3][3][3][4][5]int{}
	jd = [3]int{3, 3, 1}
	nd = [3][4]int{{0, 1, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}
)

func memset(val int) {
	for a := range dp {
		for b := range dp[a] {
			for c := range dp[a][b] {
				for d := range dp[a][b][c] {
					for e := range dp[a][b][c][d] {
						for f := range dp[a][b][c][d][e] {
							dp[a][b][c][d][e][f] = val
						}
					}
				}
			}
		}
	}
}

func solve() int {
	memset(1 << 60)
	dp[0][0][0][0][0][0] = 0
	ans := 1 << 60
	for i := 1; i < n+1; i++ {
		for a := 0; a < 2; a++ {
			for b := 0; b < 2; b++ {
				for c := 0; c < 2; c++ {
					for d := 0; d < 3; d++ {
						for e := 0; e < 4; e++ {
							v := dp[i-1][a][b][c][d][e]
							if v > 1<<60 {
								continue
							}
							if c == 0 {
								if bb[i-1] != 0 {
									continue
								}
								if bb[i] == 0 {
									dp[i][a][b][0][0][0] = min(dp[i][a][b][0][0][0], v)
								}
								if aa[i-1] == 0 {
									dp[i][a][b][1][0][1+aa[i]] =
										min(dp[i][a][b][1][0][1+aa[i]], v-i)
									if a == 0 {
										dp[i][1][b][1][1][1+1] =
											min(dp[i][1][b][1][1][1+1], v-i)
									}
								} else {
									tmp := 0
									if aa[i] == 0 {
										tmp = 1
									}
									dp[i][a][b][1][0][tmp] =
										min(dp[i][a][b][1][0][tmp], v-i)
									if a == 0 {
										dp[i][1][b][1][1][tmp] =
											min(dp[i][1][b][1][1][tmp], v-i)
									}
								}
								if b == 0 {
									dp[i][a][1][1][2][aa[i]] =
										min(dp[i][a][1][1][2][aa[i]], v-i)
								}
							} else {
								tmp := 0
								if e == jd[d] {
									tmp = e
								} else {
									if aa[i] == nd[d][e] || (d == 1 && e == 1) {
										tmp = e + 1
									} else {
										tmp = e
									}
								}
								dp[i][a][b][c][d][tmp] = min(dp[i][a][b][c][d][tmp], v)
								if bb[i] == 0 && (e >= jd[d] ||
									(d != 2 && e == jd[d]-1 && aa[i] == nd[d][e])) {
									dp[i][a][b][0][0][0] = min(dp[i][a][b][0][0][0], v+i)
								}
							}
						}
					}
				}
			}
		}
	}
	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for d := 0; d < 3; d++ {
				for e := 0; e < 4; e++ {
					ans = min(ans, dp[n][a][b][0][d][e])
				}
			}
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &s, &t)
	s = " " + s
	t = " " + t

	for i := 1; i < n+1; i++ {
		if s[i] == 'w' {
			aa[i+2] = 0
		} else {
			aa[i+2] = 1
		}
	}
	aa[n+3] = 1
	aa[n+4] = 1
	aa[n+5] = 1
	for i := 1; i < n+1; i++ {
		if t[i] == 'o' {
			bb[i+2] = 1
		} else {
			bb[i+2] = 0
		}
	}
	n += 4
	ans := solve()
	for i := 0; i < n+2; i++ {
		aa[i] ^= 1
	}
	fmt.Println(min(ans, solve()))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
