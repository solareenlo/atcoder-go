package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200003
	const K = 20
	const S = 262144

	var ava, dp [S]int
	var pw, f [K]int
	var per [N]int
	var ban [S]bool
	var comb [K][K]int

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	msk := (1 << k) - 1
	for s := 0; s <= msk; s++ {
		ava[s] = msk
	}
	pw[0] = 1
	for i := 1; i < k; i++ {
		pw[i] = pw[i-1] * 10
	}
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		var A int
		fmt.Fscan(in, &A)
		a[i] = pair{A, i}
	}
	b := make([]pair, n+1)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
		b[i] = pair{c[i], i}
	}
	sortPair(a[1:])
	sortPair(b[1:])
	for i := 1; i <= n; i++ {
		if a[i].x != b[i].x {
			fmt.Println(0)
			return
		}
		per[b[i].y] = a[i].y
	}
	for i := 1; i < n; i++ {
		S, T := 0, 0
		for j := 0; j < k; j++ {
			a := (c[i] / pw[j]) % 10
			b := (c[i+1] / pw[j]) % 10
			if a < b {
				S |= (1 << j)
			}
			if a > b {
				T |= (1 << j)
			}
		}
		ava[msk^S^T] &= (msk ^ T)
		if per[i] > per[i+1] {
			ban[msk^S] = true
		}
	}
	for i := 1; i <= msk; i <<= 1 {
		for j := 0; j <= msk; j += (i << 1) {
			for k := j; k < (j | i); k++ {
				if !ban[k] && !ban[k|i] {
					ban[k] = false
				} else {
					ban[k] = true
				}
				ava[k] &= ava[k|i]
			}
		}
	}
	fac := 1
	for i := 1; i <= k; i++ {
		f[i] = powMod(i, m)
		fac = fac * i % P
	}
	fac = powMod(fac, P-2)
	comb[0][0] = 1
	for i := 1; i <= k; i++ {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j]
			comb[i][j] = (comb[i][j] + comb[i-1][j-1]) % P
		}
	}
	for i := k; i > 0; i-- {
		tmp := 0
		for j := 0; j < i; j++ {
			if (j & 1) != 0 {
				tmp = (tmp - f[i-j]*comb[i][j]%P + P) % P
			} else {
				tmp = (tmp + f[i-j]*comb[i][j]%P) % P
			}
		}
		f[i] = tmp * fac % P
		fac = fac * i % P
	}
	dp[0] = 1
	res := 0
	for s := 0; s <= msk; s++ {
		if dp[s] == 0 {
			continue
		}
		ava[s] &= (msk ^ s)
		for i := 0; i < k; i++ {
			if ((ava[s] >> i) & 1) != 0 {
				dp[s|(1<<i)] = (dp[s|(1<<i)] + dp[s]) % P
			}
		}
		if !ban[s] {
			res = (res + f[popcount(uint32(s))]*dp[s]%P) % P
		}
	}
	fmt.Println(res)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

const P = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % P
		}
		a = a * a % P
		n /= 2
	}
	return res
}
