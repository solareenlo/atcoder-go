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

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 22
	const M = 1 << 22
	f := [N][M]int{}
	for i := 0; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < (1 << i); j++ {
			if s[j] == '1' {
				f[i][j] = 1
			}
		}
	}

	nxt := [N][M]int{}
	for i := 1; i <= n; i++ {
		for S := 0; S < (1 << i); S++ {
			p := (S >> (i - 1)) & 1
			nxt[i][S] = -1
			for j := i - 2; j >= 0; j-- {
				if (S >> j & 1) != p {
					nxt[i][S] = j + 1
					break
				}
			}
		}
	}

	cnt := [M]int{}
	len := 0
	ans := 0
	for i := 0; i <= n; i++ {
		for S := 0; S < (1 << i); S++ {
			cnt[S] = 0
		}
		for j := 0; j <= n-i; j++ {
			for S := 0; S < (1 << (i + j)); S++ {
				cnt[S>>j] += f[i+j][S]
			}
		}
		for S := (1 << i) - 1; S >= 0; S-- {
			if cnt[S] >= m {
				len = i
				ans = S
			}
		}
		for j := 1; j <= n-i; j++ {
			for S := 0; S < (1 << (i + j)); S++ {
				p := nxt[j][S&((1<<j)-1)]
				if p != -1 {
					f[i+p][((S>>j)<<p)+(S&((1<<p)-1))] += f[i+j][S]
				}
			}
		}
	}

	for i := len - 1; i >= 0; i-- {
		fmt.Fprint(out, string('0'+(ans>>i)&1))
	}
	fmt.Fprintln(out)
}
