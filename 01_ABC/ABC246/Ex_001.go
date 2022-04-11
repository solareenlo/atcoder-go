package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const mod = 998244353

type mat struct{ a [3][3]int }

var (
	A  mat
	B  mat
	C  mat
	mp = make([]mat, 401000)
	c  []string
)

func mul(a, b mat) mat {
	c := mat{}
	for k := 0; k < 3; k++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				c.a[i][j] += a.a[i][k] * b.a[k][j] % mod
				c.a[i][j] %= mod
			}
		}
	}
	return c
}

func cg(p, x int) {
	if c[x] == "0" {
		mp[p] = A
	}
	if c[x] == "1" {
		mp[p] = B
	}
	if c[x] == "?" {
		mp[p] = C
	}
}

func build(p, l, r int) {
	if l == r {
		cg(p, l)
		return
	}
	mid := (l + r) >> 1
	build(p<<1, l, mid)
	build(p<<1|1, mid+1, r)
	mp[p] = mul(mp[p<<1], mp[p<<1|1])
}

func up(p, l, r, x int) {
	if l == r {
		cg(p, l)
		return
	}
	mid := (l + r) >> 1
	if x <= mid {
		up(p<<1, l, mid, x)
	} else {
		up(p<<1|1, mid+1, r, x)
	}
	mp[p] = mul(mp[p<<1], mp[p<<1|1])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	A.a[0][0] = 1
	A.a[1][0] = 1
	A.a[2][0] = 1
	A.a[1][1] = 1
	A.a[2][2] = 1
	B.a[0][0] = 1
	B.a[0][1] = 1
	B.a[2][1] = 1
	B.a[1][1] = 1
	B.a[2][2] = 1
	C.a[0][0] = 1
	C.a[1][0] = 1
	C.a[2][0] = 1
	C.a[0][1] = 1
	C.a[1][1] = 1
	C.a[2][1] = 1
	C.a[2][2] = 1

	var n, q int
	var C string
	fmt.Fscan(in, &n, &q, &C)
	C = " " + C
	c = strings.Split(C, "")

	build(1, 1, n)

	for i := 1; i <= q; i++ {
		var x int
		var s string
		fmt.Fscan(in, &x, &s)
		c[x] = string(s[0])
		up(1, 1, n, x)
		fmt.Fprintln(out, (mp[1].a[2][0]+mp[1].a[2][1])%mod)
	}
}
