package main

import (
	"bufio"
	"fmt"
	"os"
)

const p = 924844033
const M = 524288
const N = 200001

var (
	f = [M]int{}
	w = [M | 1]int{}
	a = [N]int{}
	b = [N << 1]int{}
	c = [N << 1]int{}
	l = make([]int, M)
	r = make([]int, M)
	s = [N]int{}
	d = [N]int{}
	e = [N]int{}
	n int
	m int
	t int
)

func calc(x, y int) int {
	z := 1
	for y > 0 {
		if y&1 != 0 {
			z = z * x % p
		}
		x = x * x % p
		y >>= 1
	}
	return z
}

func NTT(o []int, u int) {
	for i := 0; i < m; i++ {
		if f[i] > i {
			t = o[f[i]]
			o[f[i]] = o[i]
			o[i] = t
		}
	}
	for j, v := 1, m&u; j < m; {
		for i := 0; i < m; i++ {
			x := o[i]
			y := w[v] * o[i+j] % p
			o[i] = (x + y) % p
			o[i+j] = (p + x - y) % p
			if (i & (j - 1)) == j-1 {
				i += j
				v = m & u
			} else {
				v += u
			}
		}
		j <<= 1
		u >>= 1
	}
}

func dfs(u, v int) {
	for i := a[u]; i > 0; i = b[i] {
		if c[i] != v {
			dfs(c[i], u)
			s[u] += s[c[i]]
			l[n-s[c[i]]]++
			l[s[c[i]]]++
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	s[n] = 1
	for i := 1; i < n; i++ {
		s[i] = 1
		var u, v int
		fmt.Fscan(in, &u, &v)
		t++
		b[t] = a[u]
		a[u] = t
		c[a[u]] = v
		t++
		b[t] = a[v]
		a[v] = t
		c[a[v]] = u
	}
	m = 1
	dfs(m, 0)
	t = 0
	for m <= n<<1 {
		m <<= 1
		t++
	}
	d[0] = 1
	d[1] = 1
	e[0] = 1
	e[1] = 1
	w[0] = 1
	w[m] = 1
	w[1] = calc(5, p>>t)
	t--
	f[1] = 1 << t
	for i := 2; i < m; i++ {
		f[i] = f[i>>1]>>1 | (i&1)<<t
		w[i] = w[i-1] * w[1] % p
	}
	for i := 2; i <= n; i++ {
		d[i] = d[i-1] * i % p
		e[i] = (p - p/i) * e[p%i] % p
	}
	for i := 2; i <= n; i++ {
		e[i] = e[i-1] * e[i] % p
	}
	for i := 0; i <= n; i++ {
		l[i] = l[i] * d[i] % p
		r[i] = e[n-i]
	}
	NTT(l, m>>1)
	NTT(r, m>>1)
	for i := 0; i < m; i++ {
		l[i] = l[i] * r[i] % p
	}
	NTT(l, -m>>1)
	t = calc(m, p-2)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, (d[n]*e[n-i]%p*e[i]%p*n%p+p-l[n+i]*t%p*e[i]%p)%p)
	}
}
