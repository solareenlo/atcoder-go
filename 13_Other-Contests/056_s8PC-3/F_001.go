package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const N = 100007

var n, m, k, t, h int
var st [N << 1]int
var d [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &m)
	h = 4*8 - clz(uint32(n))
	for i := 0; i < n; i++ {
		d[i] = -1
	}
	for i := 0; i < n; i += 2 {
		build(i + n)
	}
	for i := 1; i <= m; i++ {
		var r0, p int
		fmt.Fscan(in, &r0, &p)
		l := 1
		r := r0
		ans := 1
		for l <= r {
			mid := (l + r) / 2
			vmid := query(mid, mid)
			tmp := (r0-mid+1)*vmid - query(mid, r0)
			if tmp <= p {
				ans = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		l = ans
		vl := query(l, l)
		p -= (r0-l+1)*vl - query(l, r0)
		vnew := vl + p/(r0-l+1)
		mid := l + p%(r0-l+1) - 1
		inc(l, mid, vnew+1)
		inc(mid+1, r0, vnew)
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, query(i, i))
	}
}

func apply(p, v, k int) {
	st[p] = v * k
	if p < n {
		d[p] = v
	}
}

func build(p int) {
	k := 1
	for p > 1 {
		p >>= 1
		k <<= 1
		if d[p] == -1 {
			st[p] = st[p<<1] + st[p<<1|1]
		}
	}
}

func push(p int) {
	k := 1 << (h - 1)
	for s := h; s > 0; s, k = s-1, k>>1 {
		i := p >> s
		if d[i] == -1 {
			continue
		}
		apply(i<<1, d[i], k)
		apply(i<<1|1, d[i], k)
		d[i] = -1
	}
}

func inc(l, r, v int) {
	if l > r {
		return
	}
	l += n - 1
	r += n - 1
	push(l)
	push(r)
	l0 := l
	r0 := r
	apply(r, v, 1)
	for k := 1; l < r; l, r, k = l>>1, r>>1, k<<1 {
		if (l & 1) != 0 {
			apply(l, v, k)
			l++
		}
		if (r & 1) != 0 {
			r--
			apply(r, v, k)
		}
	}
	build(l0)
	build(r0)
}

func query(l, r int) int {
	l += n - 1
	r += n - 1
	push(l)
	push(r)
	res := st[r]
	for ; l < r; l, r = l>>1, r>>1 {
		if (l & 1) != 0 {
			res += st[l]
			l++
		}
		if (r & 1) != 0 {
			r--
			res += st[r]
		}
	}
	return res
}

func clz(x uint32) int {
	return bits.LeadingZeros32(x)
}
