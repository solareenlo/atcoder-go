package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var C, dp, F, I [1 << 17]int

	var n int
	fmt.Fscan(in, &n)
	F[0] = 1
	for i := 0; i < n; i++ {
		F[i+1] = F[i] * (i + 1) % MOD
	}
	v := int(mint(F[n]).inv())
	for i := n; i >= 1; i-- {
		I[i] = F[i-1] * v % MOD
		v = v * i % MOD
	}
	dp[0] = F[n]
	id := 0
	S := make([]int, 0)
	S = append(S, 0)
	var t string
	fmt.Fscan(in, &t)
	for i := 0; i < 2*n; i++ {
		if t[i] == '(' {
			id++
			S = append(S, id)
			dp[id] = 1
		} else {
			v := S[len(S)-1]
			S = S[:len(S)-1]
			u := S[len(S)-1]
			dp[u] = dp[u] * (dp[v] * I[C[v]+1] % MOD) % MOD
			C[u] += C[v] + 1
		}
	}
	fmt.Println(dp[0])
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
