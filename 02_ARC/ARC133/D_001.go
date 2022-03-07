package main

import "fmt"

func main() {
	var l, r, v int
	fmt.Scan(&l, &r, &v)
	l--

	const mod = 998244353
	ans := 0
	m4 := [4]int{0, 1, 3, 0}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if (m4[i] ^ m4[j]) == (v & 3) {
				dp := [66][2][2][2]int{}
				tmp0 := 0
				tmp1 := 0
				tmp2 := 0
				if (l & 3) <= i {
					tmp0 = 1
				}
				if i < j {
					tmp1 = 1
				}
				if j <= (r & 3) {
					tmp2 = 1
				}
				dp[1][tmp0][tmp1][tmp2] = 1
				for k := 2; k < 61; k++ {
					for a := 0; a < 2; a++ {
						for b := 0; b < 2; b++ {
							for c := 0; c < 2; c++ {
								for x := 0; x < 2; x++ {
									for y := 0; y < 2; y++ {
										tmp3 := x
										tmp4 := y
										if i%2 != 0 {
											tmp3 = 0
										}
										if j%2 != 0 {
											tmp4 = 0
										}
										if (tmp3 ^ tmp4) == (v >> k & 1) {
											l1 := l >> k & 1
											r1 := r >> k & 1
											tmp4 := 0
											tmp5 := 0
											tmp6 := 0
											if l1 < x || l1 == x && a != 0 {
												tmp4 = 1
											}
											if x < y || x == y && b != 0 {
												tmp5 = 1
											}
											if y < r1 || y == r1 && c != 0 {
												tmp6 = 1
											}
											dp[k][tmp4][tmp5][tmp6] += dp[k-1][a][b][c]
											dp[k][tmp4][tmp5][tmp6] %= mod
										}
									}
								}
							}
						}
					}
				}
				ans = (ans + dp[60][1][1][1]) % mod
			}
		}
	}
	fmt.Println(ans)
}
