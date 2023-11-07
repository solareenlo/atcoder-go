package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 200200

	var fa, val, tp, ans [N]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	inv := invMod(k % MOD)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &fa[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &val[i])
	}
	for i := 1; i <= n; i++ {
		tp[fa[i]] = (tp[fa[i]] + val[i]) % MOD
	}
	for i := 1; i <= n; i++ {
		val[i] = tp[i]
	}
	for k > 0 {
		if (k & 1) != 0 {
			for i := 1; i <= n; i++ {
				ans[i] = (ans[i] + val[i]) % MOD
			}
			for i := 1; i <= n; i++ {
				tp[i] = 0
			}
			for i := 1; i <= n; i++ {
				tp[fa[i]] = (tp[fa[i]] + val[i]) % MOD
			}
			for i := 1; i <= n; i++ {
				val[i] = tp[i]
			}
		}
		for i := 1; i <= n; i++ {
			tp[i] = 0
		}
		for i := 1; i <= n; i++ {
			tp[fa[i]] = (tp[fa[i]] + val[i]) % MOD
		}
		for i := 1; i <= n; i++ {
			val[i] = (val[i] + tp[i]) % MOD
		}
		for i := 1; i <= n; i++ {
			tp[i] = fa[fa[i]]
		}
		for i := 1; i <= n; i++ {
			fa[i] = tp[i]
		}
		k >>= 1
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", ans[i]*inv%MOD)
	}
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}
