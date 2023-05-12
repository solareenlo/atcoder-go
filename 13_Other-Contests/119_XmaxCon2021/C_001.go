package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const g = 3

var wk, tw [1050005]int
var ta, tb []int

var str []string
var n int
var f, jc, ny, s [250005]int

func main() {
	in := bufio.NewReader(os.Stdin)

	tw[0] = 1
	jc[0] = 1
	ny[0] = 1
	s[0] = 1
	w := powMod(g, (MOD-1)>>20)
	for i := 1; i < 1048576; i++ {
		tw[i] = tw[i-1] * w % MOD
	}
	var tmp string
	fmt.Fscan(in, &n, &tmp)
	tmp = " " + tmp + strings.Repeat(" ", 250005)
	str = strings.Split(tmp, "")
	for i := 1; i <= n+1; i++ {
		jc[i] = jc[i-1] * i % MOD
		if str[i] == "1" {
			s[i] = MOD - s[i-1]
		} else {
			s[i] = s[i-1]
		}
	}
	ny[n+1] = powMod(jc[n+1], MOD-2)
	for i := n; i > 0; i-- {
		ny[i] = ny[i+1] * (i + 1) % MOD
	}
	f[0] = 1
	ta = make([]int, 1050005)
	tb = make([]int, 1050005)
	Solve(0, n+1)
	fmt.Println(f[n+1] * jc[n+1] % MOD)
}

func Solve(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	Solve(l, mid)
	p := mid
	Len := 1
	for Len <= (mid - l + r - l) {
		Len <<= 1
	}
	for i := 0; i < Len; i++ {
		ta[i] = 0
		tb[i] = 0
	}
	for p > l && str[p] != "?" {
		p--
	}
	for i := p; i <= mid; i++ {
		if str[i] != "0" {
			ta[i-l] = f[i] * s[i] % MOD
		}
	}
	for i := 0; i <= r-l; i++ {
		tb[i] = ny[i]
	}
	DFT(ta, Len)
	DFT(tb, Len)
	for i := 0; i < Len; i++ {
		ta[i] = ta[i] * tb[i] % MOD
	}
	IDFT(ta, Len)
	for i := mid + 1; i <= r; i++ {
		if i != mid+1 && str[i-1] == "?" {
			break
		}
		f[i] = (f[i] + s[i-1]*ta[i-l]) % MOD
	}
	Solve(mid+1, r)
}

func DFT(a []int, n int) {
	for i := n >> 1; i > 0; i >>= 1 {
		tmp := 524288 / i
		for j := 0; j < i; j++ {
			wk[j] = tw[tmp*j]
		}
		for j := 0; j < n; j += (i << 1) {
			for k := 0; k < i; k++ {
				x := a[j+k]
				y := a[i+j+k]
				z := x
				x = (x + y) % MOD
				a[j+k] = x
				z = (z - y + MOD) % MOD
				a[i+j+k] = z * wk[k] % MOD
			}
		}
	}
}

func IDFT(a []int, n int) {
	for i := 1; i < n; i <<= 1 {
		tmp := 524288 / i
		for j := 0; j < i; j++ {
			if j != 0 {
				wk[j] = tw[tmp*(2*i-j)]
			} else {
				wk[j] = tw[0]
			}
		}
		for j := 0; j < n; j += (i << 1) {
			for k := 0; k < i; k++ {
				x := a[j+k]
				y := a[i+j+k] * wk[k] % MOD
				z := x
				x = (x + y) % MOD
				a[j+k] = x
				z = (z - y + MOD) % MOD
				a[i+j+k] = z
			}
		}
	}
	inv := powMod(n, MOD-2)
	for i := 0; i < n; i++ {
		a[i] = a[i] * inv % MOD
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
