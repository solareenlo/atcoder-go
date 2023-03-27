package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	x, y, next int
}

var Len int
var du, last, sum, cnt, tot, f, w [100010]int
var a [200010]node
var p [100010][]int
var v [100010]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		du[x]++
		du[y]++
		ins(x, y)
		ins(y, x)
	}
	for i := 1; i <= Len; i += 2 {
		x := a[i].x
		y := a[i].y
		if pd(x, y) == true {
			x, y = y, x
		}
		p[x] = append(p[x], y)
		sum[x] += du[y] - 1
		sum[y] += du[x] - 1
	}
	for i := 1; i <= n; i++ {
		x := i
		for j := 0; j < len(p[x]); j++ {
			v[p[x][j]] = true
		}
		for j := 0; j < len(p[x]); j++ {
			y := p[x][j]
			for k := 0; k < len(p[y]); k++ {
				z := p[y][k]
				if v[z] == true {
					cnt[x]++
					cnt[y]++
					cnt[z]++
				}
			}
		}
		for j := 0; j < len(p[x]); j++ {
			v[p[x][j]] = false
		}
	}
	for i := 1; i <= n; i++ {
		x := i
		for k := last[x]; k > 0; k = a[k].next {
			f[x] += sum[a[k].y]
		}
		f[x] -= du[x] * (du[x] - 1)
		f[x] -= cnt[x] * 2
	}
	for i := 1; i <= n; i++ {
		x := i
		for k := last[x]; k > 0; k = a[k].next {
			y := a[k].y
			for j := 0; j < len(p[y]); j++ {
				z := p[y][j]
				if pd(x, z) == true {
					continue
				}
				tot[x] += w[z]
				tot[y] += w[z]
				tot[z] += w[z]
				w[z]++
			}
		}
		for k := last[x]; k > 0; k = a[k].next {
			y := a[k].y
			for j := 0; j < len(p[y]); j++ {
				z := p[y][j]
				if pd(x, z) == true {
					continue
				}
				w[z]--
				tot[y] += w[z]
			}
		}
	}
	for i := 1; i <= n; i++ {
		x := i
		ans := 0
		for k := last[x]; k > 0; k = a[k].next {
			ans += f[a[k].y]
		}
		ans -= 2 * tot[x]
		ans -= 2 * cnt[x] * (du[x] - 2)
		ans -= sum[x]*(du[x]-1) - 2*cnt[x]
		fmt.Fprintln(out, ans)
	}
}

func ins(x, y int) {
	Len++
	a[Len].x = x
	a[Len].y = y
	a[Len].next = last[x]
	last[x] = Len
}

func pd(x, y int) bool {
	if du[x] > du[y] || (du[x] == du[y] && x >= y) {
		return true
	} else {
		return false
	}
}
