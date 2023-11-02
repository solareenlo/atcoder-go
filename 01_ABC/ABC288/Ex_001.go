package main

import "fmt"

const N = 205
const M = 35
const MOD = 998244353

var n, m, x int
var f, g [][N]int
var inv [N]int
var mp [N]map[int]int

func nCr(n, r int) int {
	if _, ok := mp[r][n]; ok {
		return mp[r][n]
	}
	ret := inv[r]
	for i := n - r + 1; i <= n; i++ {
		ret = ret * i % MOD
	}
	mp[r][n] = ret
	return mp[r][n]
}

func powMod(a, x int) int {
	rt := 1
	for x > 0 {
		if (x & 1) != 0 {
			rt = rt * a % MOD
		}
		a = a * a % MOD
		x >>= 1
	}
	return rt
}

func solve(_m, _x int, f [][N]int) {
	x := _x
	m := _m
	for i := 0; i < n+1; i++ {
		f[0][i] = 1
	}
	for t := 1; t <= 30; t++ {
		tx := x & (1 << (t - 1))
		if tx != 0 {
			tx = 1
		}
		tm := m & (1 << (t - 1))
		s := (m & ((1 << (t - 1)) - 1)) + 1
		b := (1 << (t - 1))
		if tm == 0 {
			if tx == 0 {
				for i := 0; i < n+1; i++ {
					f[t][i] = f[t-1][i]
				}
			}
			continue
		}
		I := powMod(b-1, MOD-2)
		for i := 0; i < n+1; i++ {
			for j := 0; j < i+1; j++ {
				if (j & 1) == tx {
					if i == j {
						f[t][i] = (f[t][i] + f[t-1][i]) % MOD
					} else {
						tmp0 := nCr(s+j-1, j)
						tmp1 := nCr(b+i-j-1, i-j)
						f[t][i] = (f[t][i] + ((tmp0-f[t-1][j])*(tmp1-g[t-1][i-j])%MOD)*I) % MOD
						f[t][i] = (f[t][i] + f[t-1][j]*g[t-1][i-j]%MOD) % MOD
					}
				}
			}
		}
	}
}

func main() {
	f = make([][N]int, M)
	g = make([][N]int, M)
	for i := range mp {
		mp[i] = make(map[int]int)
	}
	fmt.Scan(&n, &m, &x)
	inv[0] = 1
	for i := 1; i <= n; i++ {
		inv[i] = inv[i-1] * powMod(i, MOD-2) % MOD
	}
	solve((1<<30)-1, 0, g)
	solve(m, x, f)
	fmt.Println((f[30][n] + MOD) % MOD)
}
