package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const mod = 998244353

var (
	n   int
	x   = [25]int{}
	L   = [25]int{}
	R   = [25]int{}
	ans int
)

func ksm(x, y int) int {
	res := 1
	for y > 0 {
		if y&1 != 0 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}

type node struct{ x, y, p, q int }

var (
	l    = [50]int{}
	r    = [50]int{}
	d0   = [50]int{}
	d1   = [50]int{}
	f    = [50][50]int{}
	g    = [50]int{}
	p    = [50]int{}
	ivjc = [25]int{}
	tot  int
	a    = [2000]node{}
)

func update(x *int, y int) {
	*x += y
	if *x >= mod {
		*x -= mod
	}
}

func X(i int) int {
	if i > n {
		return R[i-n]
	}
	return L[i]
}

func I(i int) int {
	if i > n {
		return i - n
	}
	return i
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	ivjc[0] = 1
	for i := 1; i <= n; i++ {
		ivjc[i] = ivjc[i-1] * ksm(i, mod-2) % mod
	}
	for i := 0; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}
	for i := 1; i <= n; i++ {
		L[i] = x[i-1]
		R[i] = x[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			for k1 := 0; k1 < 2; k1++ {
				for k2 := 0; k2 < 2; k2++ {
					tot++
					a[tot].x = j + k1*n
					a[tot].y = i + k2*n
					tmp1 := L[i]
					if k2 != 0 {
						tmp1 = R[i]
					}
					tmp2 := L[j]
					if k1 != 0 {
						tmp2 = R[j]
					}
					a[tot].p = tmp1 - tmp2
					a[tot].q = i - j
				}
			}
		}
	}
	tmp := a[1 : tot+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].p*tmp[j].q < tmp[i].q*tmp[j].p
	})
	for i := 1; i <= n; i++ {
		p[0]++
		p[p[0]] = i
		p[0]++
		p[p[0]] = i + n
	}
	for now := 1; now < tot; now++ {
		flag := 0
		for i := 1; i <= n*2; i++ {
			if p[i] == a[now].y {
				for j := i + 1; j <= n*2; j++ {
					if p[j] == a[now].x {
						flag = 1
					}
				}
				if flag != 0 {
					break
				}
				for j := i; j < n*2; j++ {
					p[j] = p[j+1]
				}
				break
			}
		}
		if flag == 0 {
			for i := 1; i < n*2; i++ {
				if p[i] == a[now].x {
					for j := n * 2; j > i; j-- {
						p[j] = p[j-1]
					}
					p[i] = a[now].y
					break
				}
			}
		}
		if a[now].p*a[now+1].q == a[now+1].p*a[now].q {
			continue
		}
		for i := 1; i <= n*2; i++ {
			if p[i] <= n {
				l[p[i]] = i
			} else {
				r[p[i]-n] = i
			}
		}
		for i := 2; i <= n; i++ {
			l[i] = max(l[i], l[i-1])
		}
		for i := n - 1; i >= 1; i-- {
			r[i] = min(r[i], r[i+1])
		}
		flag = 1
		for i := 1; i <= n; i++ {
			if l[i] >= r[i] {
				flag = 0
				break
			}
		}
		if flag == 0 {
			break
		}
		for i := range d0 {
			d0[i] = 0
		}
		for i := range d1 {
			d1[i] = 0
		}
		for i := 1; i <= n; i++ {
			d0[l[i]] = max(d0[l[i]], i)
			d1[r[i]] = max(d1[r[i]], i)
		}
		for i := range f {
			for j := range f[i] {
				f[i][j] = 0
			}
		}
		f[0][0] = 1
		st := 1
		ed := 0
		for i := 1; i < n*2; i++ {
			ed = max(ed, d0[i])
			st = max(st, d1[i]+1)
			if ed < st {
				continue
			}
			for j := ed - 1; j >= st-1; j-- {
				for i := range g {
					g[i] = f[j][i]
				}
				u := I(p[i]) - I(p[i+1])
				v := X(p[i+1]) - X(p[i])
				u = (u%mod + mod) % mod
				v = (v%mod + mod) % mod
				for k := 1; k <= ed-j; k++ {
					for t := j + k; t >= 1; t-- {
						g[t] = (g[t]*v + g[t-1]*u) % mod
					}
					g[0] = g[0] * v % mod
					for t := 0; t <= j+k; t++ {
						update(&f[j+k][t], g[t]*ivjc[k]%mod)
					}
				}
			}
		}
		X := a[now].p * ksm(a[now].q, mod-2) % mod
		Y := a[now+1].p * ksm(a[now+1].q, mod-2) % mod
		for i := 0; i <= n; i++ {
			update(&ans, f[n][i]*ksm(i+1, mod-2)%mod*
				(ksm(Y, i+1)-ksm(X, i+1)+mod)%mod)
		}
	}

	for i := 1; i <= n; i++ {
		ans = ans * ksm(x[i]-x[i-1], mod-2) % mod
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
