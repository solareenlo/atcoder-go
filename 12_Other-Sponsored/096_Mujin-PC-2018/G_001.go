package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func count(Len int) int {
	if Len >= 0 {
		return (Len + 1) * (Len + 2) % MOD * (Len + 3) % MOD * 166374059 % MOD
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &Q)
	for q := 0; q < Q; q++ {
		var xa, ya, xb, yb, xc, yc, k int
		fmt.Fscan(in, &xa, &ya, &xb, &yb, &xc, &yc, &k)
		a := yb*xc - xb*yc
		b := yc*xa - xc*ya
		c := ya*xb - xa*yb
		g := gcd(a, gcd(b, c))
		a /= g
		b /= g
		c /= g
		if a+b+c > 0 {
			a *= -1
			b *= -1
			c *= -1
		}
		fmt.Fprintln(out, (count(k)+MOD-count(k+a+b+c-max(a, 0)-max(b, 0)-max(c, 0)))%MOD)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
