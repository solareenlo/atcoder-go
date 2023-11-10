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

	const N = 200005
	const V = 200000

	var a, cnt, mu, b [N]int
	var f [N][2]bool

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}
	mu[1] = 1
	for i := 1; i <= V; i++ {
		for j := i << 1; j <= V; j += i {
			mu[j] -= mu[i]
			cnt[i] += cnt[j]
		}
	}
	for i := 2; i <= V; i++ {
		if !f[i][1] && f[i][0] {
			f[i][1] = false
		} else {
			f[i][1] = true
		}
		for j := i; j <= V; j += i {
			b[j] = 0
		}
		for d, tmp := 1, i; tmp <= V; d, tmp = d+1, tmp+i {
			for j := tmp; j <= V; j += tmp {
				b[j] += mu[d] * cnt[tmp]
			}
		}
		for j := i << 1; j <= V; j += i {
			for x := 0; x < 2 && b[j] != 0; x++ {
				if !f[j][x] && f[i][(x^cnt[i]^cnt[j]^1)&1] {
					f[j][x] = false
				} else {
					f[j][x] = true
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		if f[a[i]][(cnt[a[i]]^1)&1] {
			fmt.Fprintln(out, "Aoki")
		} else {
			fmt.Fprintln(out, "Takahashi")
		}
	}
}
