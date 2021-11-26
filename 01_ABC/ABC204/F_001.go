package main

import "fmt"

var (
	mod     = 998244353
	h, w, t int
	res     mat
	ans     mat
	base    mat
)

type mat struct{ g [70][70]int }

func mult(a, b mat) mat {
	var c mat
	for k := 0; k < 64; k++ {
		for j := 0; j < 64; j++ {
			for i := 0; i < 64; i++ {
				c.g[i][j] = (c.g[i][j] + a.g[i][k]*b.g[k][j]%mod) % mod
			}
		}
	}
	return c
}

func qpow(pw int) {
	for pw > 0 {
		if pw&1 != 0 {
			ans = mult(ans, base)
		}
		base = mult(base, base)
		pw >>= 1
	}
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	res.g[0][0] = 1
	t = 1 << h
	for i := 0; i < t; i++ {
		for j := 0; j < t; j++ {
			for k := 0; k < t; k++ {
				if i+j+k+(k<<1) == (i|j|k|(k<<1)) && (i+j+k+(k<<1) < t) {
					base.g[i][j]++
				}
			}
		}
	}

	ans = base
	qpow(w - 1)
	res = mult(res, ans)
	fmt.Println(res.g[0][0])
}
