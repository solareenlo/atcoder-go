package main

import (
	"fmt"
	"math/bits"
)

func main() {
	const N = 105
	const M = 270
	const MOD = 998244353

	var f [N][N][M]int

	var n, m int
	fmt.Scan(&n, &m)
	S := (1 << (2*m - 2)) - 1
	f[0][0][S] = 1
	l, r := 0, 0
	for i := 0; i <= 2*m-2-1; i++ {
		l |= 1 << i
		i++
	}
	for i := 1; i <= 2*m-2-1; i++ {
		r |= 1 << i
		i++
	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= i; j++ {
			for s := 0; s <= S; s++ {
				if f[i][j][s] == 0 {
					continue
				}
				u := i - j - popcount(uint32(S^((s&l)|r)))
				v := i - j - popcount(uint32(S^((s&r)|l)))
				w := f[i][j][s]
				z := s << 2
				f[i+1][j][z&S] = (f[i+1][j][z&S] + w) % MOD
				f[i+1][j+1][(z|1)&S] = (f[i+1][j+1][(z|1)&S] + v*w) % MOD
				f[i+1][j+1][(z|2)&S] = (f[i+1][j+1][(z|2)&S] + u*w) % MOD
				f[i+1][j+2][(z|3)&S] = (f[i+1][j+2][(z|3)&S] + u*v*w) % MOD
			}
		}
	}
	fmt.Println(f[n][n][S])
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
