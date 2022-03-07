package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005
	const mod = 998244353

	pw := make([]int, N)
	pw[0] = 1
	for i := 1; i < N; i++ {
		pw[i] = 3 * pw[i-1] % mod
	}

	var s string
	var m int
	fmt.Fscan(in, &s, &m)
	n := len(s)
	s = " " + s

	ts := [2][N]int{}
	ts[0][0] = 1
	a := make([]int, N)
	b := make([]int, N)
	f := [2][2][N]int{}
	g := [2][N]int{}
	t := 0
	las := -1
	for i := 1; i+2 <= n; i++ {
		if s[i] == 'A' && s[i+1] == 'R' && s[i+2] == 'C' {
			at := 0
			bt := 0
			l := i - 1
			r := i + 3
			for l > 0 {
				if s[l] == 'A' {
					at++
					a[at] = 1
					l--
				} else if l > 1 && s[l-1] == 'A' && s[l] == 'R' {
					at++
					a[at] = 2
					l -= 2
				} else {
					break
				}
			}
			for r <= n {
				if s[r] == 'C' {
					bt++
					b[bt] = 1
					r++
				} else if r < n && s[r] == 'R' && s[r+1] == 'C' {
					bt++
					b[bt] = 2
					r += 2
				} else {
					break
				}
			}
			ft := at + bt + 2
			for j := 0; j <= ft; j++ {
				for fl0 := 0; fl0 < 2; fl0++ {
					for fl1 := 0; fl1 < 2; fl1++ {
						f[fl0][fl1][j] = 0
					}
				}
			}
			f[0][0][0] = 1
			b[bt+1] = 1
			A := 0
			for j := 0; j <= at; j++ {
				B := 0
				for k := 0; k <= bt+1; k++ {
					if j == 0 && k == 0 {
						f[0][0][j+k+1] = pw[3] - 1
						if j == at {
							f[1][0][j+k+1] = pw[3]
						}
						continue
					}
					q := pw[A+B+3]
					q2 := q
					if a[j] != 0 {
						q = q * (pw[a[j]] - 1) % mod
						q2 = q2 * pw[a[j]] % mod
					}
					if b[k] != 0 {
						q = q * (pw[b[k]] - 1) % mod
						q2 = q2 * (pw[b[k]] - 1) % mod
					}
					tmp := 0
					if k > bt {
						tmp = 1
					}
					f[0][tmp][j+k+1] += q
					f[0][tmp][j+k+1] %= mod
					if j == at {
						f[1][tmp][j+k+1] += q2
						f[1][tmp][j+k+1] %= mod
					}
					B += b[k]
				}
				A += a[j]
			}
			if las != l || s[las] != 'R' {
				for j := 0; j <= t; j++ {
					ts[1][j] = 0
				}
			}
			for j := 0; j <= t+ft; j++ {
				g[0][j] = 0
				g[1][j] = 0
			}
			for fl0 := 0; fl0 < 2; fl0++ {
				for fl1 := 0; fl1 < 2; fl1++ {
					for j := 0; j <= t; j++ {
						for k := 0; k <= ft; k++ {
							g[fl1][j+k] = (g[fl1][j+k] + ts[fl0][j]*f[fl0][fl1][k]) % mod
						}
					}
				}
			}
			t += ft
			for j := 0; j <= t; j++ {
				ts[0][j] = g[0][j]
				ts[1][j] = g[1][j]
			}
			las = r
		}
	}

	ans := 0
	for j := 0; j <= m; j++ {
		ans += ts[0][j]
		ans %= mod
	}
	fmt.Println(ans)
}
