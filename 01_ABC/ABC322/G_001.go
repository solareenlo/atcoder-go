package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var n, x int
	fmt.Fscan(in, &n, &x)

	ans := 0
	for d := 1; d < 10; d++ {
		xx := x
		if x%d != 0 {
			continue
		}
		x /= d
		l := max(x+d+1, d+1)
		r := n
		for A := l; A <= min(l+20, r); A++ {
			ans += min(10, A-x)
			ans %= MOD
		}
		l += 21
		ans += 10 * max(0, r-l+1) % MOD
		ans %= MOD
		x = xx
	}

	for a := 2; a <= n; a++ {
		if a*a-(a-1)*(a-1) > x {
			break
		}
		for b := a - 1; b >= 2; b-- {
			xx := x
			A := 1
			B := 1
			for A-B <= x {
				A *= a
				B *= b
			}
			if A == a*a {
				break
			}
			for xx > 0 && A-B > 0 {
				d := xx / (A - B)
				d = min(min(d, 9), min(a-1, b-1))
				xx -= d * (A - B)
				A /= a
				B /= b
			}
			if xx == 0 {
				ans += min(10, a, b)
				ans %= MOD
			}
		}
	}
	fmt.Println(ans)
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
