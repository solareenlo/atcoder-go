package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAT = 4
const mod = 1000000007

type Mat struct {
	v [4][4]int
}

func mulmat(a, b Mat, n int) Mat {
	var r Mat
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r.v[i][j] = 0
			for k := 0; k < n; k++ {
				r.v[i][j] += (a.v[i][k] * b.v[k][j]) % mod
			}
			r.v[i][j] %= mod
		}
	}
	return r
}

func powmat(p int, a Mat, n int) Mat {
	var r Mat
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r.v[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		r.v[i][i] = 1
	}
	for p > 0 {
		if p%2 != 0 {
			r = mulmat(r, a, n)
		}
		a = mulmat(a, a, n)
		p >>= 1
	}
	return r
}

const NV = 1 << 17

var (
	val  = [NV * 2]Mat{}
	S    = [20050]int{}
	T    = [20050]int{}
	mask = [100001]int{}
	XM   = map[int]int{}
	XM2  = make([]int, 0)
)

func getmat(h, mask int) Mat {
	var r Mat
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r.v[i][j] = 0
		}
	}
	r.v[0][0] = 1
	if h == 1 {
		if mask == 0 {
			r.v[0][1] = 1
			r.v[1][0] = 1
		}
	} else {
		if mask == 0 {
			r.v[0][0] = 2
			r.v[3][0] = 1
			r.v[2][1] = 1
			r.v[1][2] = 1
			r.v[0][3] = 1
		}
		if mask == 1 || mask == 0 {
			r.v[2][0] = 1
			r.v[0][2] = 1
		}
		if mask == 2 || mask == 0 {
			r.v[1][0] = 1
			r.v[0][1] = 1
		}
	}
	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var H, W, n int
	fmt.Fscan(in, &H, &W, &n)

	S := make([]int, n)
	T := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &S[i], &T[i])
	}
	for i := 0; i < n; i++ {
		S[i]--
		T[i]--
		XM[T[i]] = 0
		XM[T[i]+1] = 0
	}
	XM[0] = 0
	XM[W] = 0
	keys := make([]int, 0, len(XM))
	for k := range XM {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		XM[k] = len(XM2)
		XM2 = append(XM2, k)
	}

	for i := 0; i < NV; i++ {
		val[i+NV].v[0][0] = 1
	}

	b := getmat(H, 0)

	for i := 0; i < len(XM)-1; i++ {
		val[i+NV] = powmat(XM2[i+1]-XM2[i], b, MAT)
	}

	for i := NV - 1; i >= 0; i-- {
		if i == 0 {
			continue
		}
		val[i] = mulmat(val[i*2], val[i*2+1], MAT)
	}

	for i := 0; i < n; i++ {
		x := XM[T[i]]
		mask[x] ^= 1 << S[i]
		val[x+NV] = getmat(H, mask[x])
		x += NV
		for x > 1 {
			x >>= 1
			val[x] = mulmat(val[x*2], val[x*2+1], MAT)
		}
		fmt.Fprintln(out, val[1].v[0][0])
	}
}
