package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353
const o = 2000010

var fac, inv [o]int

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 510

	var n, m, p int
	fmt.Scan(&n, &m, &p)
	inv[1] = 1
	m += n
	m_ := m / p * p

	var f [N][N]int
	f[0][0] = 1
	for i := m_; i < m; i++ {
		for j := i - m_; j+1 > 0; j-- {
			for k := 0; k < p; k++ {
				if f[j][k] != 0 {
					f[j+1][(k+i)%p] = (f[j+1][(k+i)%p] + f[j][k]) % MOD
				}
			}
		}
	}
	for i := 2; i <= m || i <= p; i++ {
		inv[i] = MOD - MOD/i*inv[MOD%i]%MOD
	}
	var coef [N][N]int
	for i := 0; i < p; i++ {
		for j := 1; j <= p; j++ {
			if p%j == 0 && i*j%p == 0 {
				coef[i][j] = inv[j]
			} else {
				coef[i][j] = 0
			}
		}
		for j := p; j > 0; j-- {
			if p%j == 0 {
				for k := j * 2; k <= p; k += j {
					coef[i][j] = (coef[i][j] + MOD - coef[i][k]) % MOD
				}
			}
		}
	}
	fac[0] = 1
	inv[0] = 1
	for i := 1; i <= m; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = inv[i-1] * inv[i] % MOD
	}
	v := make([]int, N)
	ans := make([]int, N)
	for i := 0; i <= m-m_ && i <= n; i++ {
		for j := 0; j < p; j++ {
			v[j] = 0
		}
		for j := 1; j <= p; j++ {
			if p%j == 0 && (n-i)%(p/j) == 0 {
				t := (n - i) / (p / j)
				tmp := 1
				if ((t + n - i) & 1) != 0 {
					tmp = MOD - 1
				}
				t = nCr(j*m_/p, t) * tmp % MOD
				for k := 0; k < p; k++ {
					v[k] = (v[k] + t*coef[k][j]) % MOD
				}
			}
		}
		for j := 0; j < p; j++ {
			if f[i][j] != 0 {
				for k := 0; k < p; k++ {
					ans[(j+k)%p] = (ans[(j+k)%p] + f[i][j]*v[k]%MOD) % MOD
				}
			}
		}
	}
	delt := 0
	for i := 0; i < n; i++ {
		delt = (delt + i) % p
	}
	for i := 0; i < p; i++ {
		fmt.Fprintf(out, "%d ", ans[(delt+i)%p])
	}
}

func nCr(x, y int) int {
	if x < y {
		return 0
	}
	return fac[x] * inv[y] % MOD * inv[x-y] % MOD
}
