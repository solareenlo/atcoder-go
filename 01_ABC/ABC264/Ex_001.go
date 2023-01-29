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

	const mod = 998244353
	const N = 300030
	const M = 21

	var n int
	fmt.Fscan(in, &n)
	var p [N]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &p[i])
		p[i]--
	}

	res := 1
	var f [N][M]int
	var s [N][M]int
	var d [N]int
	fmt.Fprintln(out, 1)
	for i := 1; i < n; i++ {
		d[i] = d[p[i]] + 1
		t := d[i]
		if t < M {
			v := 1
			f[i][t] = 1
			for x := i; x > 0; x = p[x] {
				y := p[x]
				ts := v
				v *= ((s[y][t]-f[x][t]+mod)%mod + ts) % mod
				v %= mod
				s[y][t] += ts
				s[y][t] %= mod
				f[y][t] += v
				f[y][t] %= mod
			}
			res += v
			res %= mod
		}
		fmt.Fprintln(out, res)
	}

}
