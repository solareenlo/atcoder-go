package main

import (
	"bufio"
	"fmt"
	"os"
)

/* https://en.wikipedia.org/wiki/Stern%E2%80%93Brocot_tree */

const N = 100000
const LN = 17 /* LN = ceil(log2(N)) */
const N_ = (N*(LN+1) + 1)
const MOD = 998244353

var xx_, yy_ *[]int
var X uint32

func main() {
	in := bufio.NewReader(os.Stdin)

	X = 1675165015

	var n int
	fmt.Fscan(in, &n)
	xx := make([]int, N)
	yy := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xx[i], &yy[i])
	}

	ii = make([]int, N)
	dd = make([]int, N+2)
	tt = make([]int, N+2)
	ans := 0
	A = 1
	for i, j := -1, 0; j <= n; j++ {
		if j == n || gcd(xx[j], yy[j]) != 1 {
			tmpX := xx[i+1:]
			tmpY := yy[i+1:]
			ans += solve(&tmpX, &tmpY, j-i-1)
			ans %= MOD
			i = j
		}
	}

	fmt.Println(ans)
}

var ii, dd, tt []int

func solve(xx, yy *[]int, n int) int {
	n_ := 0
	for i := 0; i < n; i++ {
		if (*xx)[i] != 0 && (*yy)[i] != 0 {
			ii[n_] = i
			n_++
		}
	}
	xx_, yy_ = xx, yy
	sort_(ii, 0, n_)
	cnt := 0
	dd[cnt] = -1
	tt[cnt] = 0
	cnt++
	dd[cnt] = 0
	tt[cnt] = 0
	cnt++
	ans := 0
	for i := 0; i < n_; i++ {
		d := lcp(ii[i], ii[i])
		p := 0
		if i != 0 {
			p = lcp(ii[i-1], ii[i])
		}
		t := 0
		var c int
		for {
			t = merge(t, tt[cnt-1])
			if t == 0 {
				c = 0
			} else {
				c = (choose2(n+1) - cc[t] + MOD) % MOD
				c = (c - choose2(aa[t]+1) + MOD) % MOD
				c = (c - choose2(n-bb[t]) + MOD) % MOD
			}
			if p < dd[cnt-2] {
				ans += c * (dd[cnt-1] - dd[cnt-2] + MOD) % MOD
				ans %= MOD
				cnt--
			} else {
				ans += c * (dd[cnt-1] - p + MOD) % MOD
				ans %= MOD
				dd[cnt-1] = p
				tt[cnt-1] = t
				dd[cnt] = d
				tt[cnt] = node(0, n, ii[i])
				cnt++
				break
			}
		}
	}
	t := 0
	for cnt >= 2 {
		var c int
		t = merge(t, tt[cnt-1])
		if t == 0 {
			c = 0
		} else {
			c = (choose2(n+1) - cc[t] + MOD) % MOD
			c = (c - choose2(aa[t]+1) + MOD) % MOD
			c = (c - choose2(n-bb[t]) + MOD) % MOD
		}
		ans += c * (dd[cnt-1] - dd[cnt-2] + MOD) % MOD
		ans %= MOD
		cnt--
	}
	return ans
}

func sort_(ii []int, l, r int) {
	for l < r {
		i := l
		j := l
		k := r
		tmp := int(rand_())
		i_ := ii[l+(tmp%(r-l))]
		for j < k {
			c := sgn(cross(ii[j], i_))
			if c == 0 {
				j++
			} else if c < 0 {
				ii[i], ii[j] = ii[j], ii[i]
				i++
				j++
			} else {
				k--
				ii[j], ii[k] = ii[k], ii[j]
			}
		}
		sort_(ii, l, i)
		l = k
	}
}

func rand_() uint32 {
	X *= 3
	return X >> 1
}

func cross(i, j int) int {
	return (*xx_)[i]*(*yy_)[j] - (*xx_)[j]*(*yy_)[i]
}

var A int

func node(l, r, i int) int {
	t := A
	A++
	if r-l > 1 {
		m := (l + r) / 2
		if i < m {
			ll[t] = node(l, m, i)
		} else {
			rr[t] = node(m, r, i)
		}
	}
	aa[t] = i
	bb[t] = i
	cc[t] = 0
	return t
}

func lcp(i, j int) int {
	x1 := (*xx_)[i]
	y1 := (*yy_)[i]
	x2 := (*xx_)[j]
	y2 := (*yy_)[j]
	d := 0
	for {
		if sgn(x1-y1) != sgn(x2-y2) || x1 == y1 {
			return d
		}
		if x1 > y1 {
			d += min((x1-1)/y1, (x2-1)/y2)
			if (x1-1)/y1 != (x2-1)/y2 {
				return d
			}
			x1 = (x1-1)%y1 + 1
			x2 = (x2-1)%y2 + 1
		} else {
			d += min((y1-1)/x1, (y2-1)/x2)
			if (y1-1)/x1 != (y2-1)/x2 {
				return d
			}
			y1 = (y1-1)%x1 + 1
			y2 = (y2-1)%x2 + 1
		}
	}
}

var ll, rr, aa, bb, cc [N_]int

func merge(u, v int) int {
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}
	l := merge(ll[u], ll[v])
	r := merge(rr[u], rr[v])
	ll[u] = l
	rr[u] = r
	tmp := r
	if l != 0 {
		tmp = l
	}
	aa[u] = aa[tmp]
	tmp = l
	if r != 0 {
		tmp = r
	}
	bb[u] = bb[tmp]
	tmp = 0
	if l != 0 && r != 0 {
		tmp = choose2(aa[r] - bb[l])
	}
	cc[u] = (cc[l] + cc[r] + tmp) % MOD
	return u
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func choose2(a int) int {
	return a * (a - 1) / 2 % MOD
}

func sgn(a int) int {
	if a == 0 {
		return 0
	}
	if a > 0 {
		return 1
	}
	return -1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
