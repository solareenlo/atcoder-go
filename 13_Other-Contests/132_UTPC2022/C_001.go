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
	M := 512
	ans := (c2(m+1) * ex(m, n-1) % MOD) * n % MOD
	pws := make([]int, M*2)
	for i := 0; i < M*2; i++ {
		pws[i] = int(mint(i - M).pow(n))
	}
	for x := 1; x < M; x++ {
		d := make([]int, M)
		for a := x; ; a = x & (a - 1) {
			e := make([]int, M)
			for i := 1; i <= m; i++ {
				if (i & x) <= a {
					e[i] += 1
				}
			}
			ada(e)
			for i := 0; i < M; i++ {
				e[i] = pws[(e[i]+M)%MOD]
			}
			dad(e)
			d[a] = e[x]
			if a == 0 {
				break
			}
		}
		for i := 0; i < M; i++ {
			if (i & x) != i {
				d[i] = d[i-1]
			}
		}
		ans += d[len(d)-1] * (x + 1)
		for i := 1; i <= M-1; i++ {
			ans = (ans - ((d[len(d)-1]-d[i-1]+MOD)%MOD)*2%MOD + MOD) % MOD
		}
	}
	fmt.Println(ans)
}

func c2(n int) int { return n * (n - 1) >> 1 }

func ex(x, t int) int { return int(mint(x).pow(t)) }

func ada(d []int) {
	n := len(d)
	for i := n; i >= 1; i >>= 1 {
		for j := 0; j < n; j++ {
			if (j & i) != 0 {
				x := d[j^i]
				y := d[j]
				d[j^i] = (x + y) % MOD
				d[j] = (x - y + MOD) % MOD
			}
		}
	}
}

func dad(d []int) {
	n := len(d)
	i2 := ex(2, MOD-2)
	for i := 1; i <= n; i <<= 1 {
		for j := 0; j < n; j++ {
			if (j & i) != 0 {
				x := d[j^i]
				y := d[j]
				d[j^i] = (x + y) * i2 % MOD
				d[j] = ((x - y + MOD) % MOD) * i2 % MOD
			}
		}
	}
}

type mint int

func (m mint) pow(p int) mint {
	return powMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const MOD = 998244353

func powMod(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return powMod(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
