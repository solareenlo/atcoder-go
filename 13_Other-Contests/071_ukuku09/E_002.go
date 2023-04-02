package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROOT = 5
const MOD = 924844033
const MAX_N = 200001

var fac, finv, inv [MAX_N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < MAX_N; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}

	var N int
	fmt.Fscan(in, &N)
	a := make([]int, 2*N)

	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x-1] = 1
	}

	tmp := 0
	polys := make([][]int, 0)
	for i := 0; i < 2*N-1; i++ {
		if a[i] != a[i+1] {
			tmp++
		} else {
			if tmp == 0 {
				continue
			}
			v := make([]int, 0)
			for m := 0; tmp-m+1 >= m; m++ {
				v = append(v, mul(mul(fac[tmp-m+1], finv[m]), finv[tmp-2*m+1]))
			}
			polys = append(polys, v)
			tmp = 0
		}
	}
	if tmp != 0 {
		v := make([]int, 0)
		for m := 0; tmp-m+1 >= m; m++ {
			v = append(v, mul(mul(fac[tmp-m+1], finv[m]), finv[tmp-2*m+1]))
		}
		polys = append(polys, v)
	}

	prod := allProd(polys, 0, len(polys))
	ans := 0
	sgn := 1
	for i := 0; i < len(prod); i++ {
		ans = (ans + (MOD + sgn*((prod[i]*fac[N-i])%MOD))) % MOD
		sgn *= -1
	}
	fmt.Println(ans)
}

func mul(x, y int) int {
	return x * y % MOD
}

func allProd(polys [][]int, l, r int) []int {
	if l >= r {
		return []int{1, 1}
	}
	if l+1 == r {
		return polys[l]
	}
	vl := allProd(polys, l, (l+r)/2)
	vr := allProd(polys, (l+r)/2, r)
	return convolute(vl, vr)
}

func convolute(a, b []int) []int {
	size := len(a) + len(b) - 1
	t := 1
	for t < size {
		t <<= 1
	}
	A := make([]int, t)
	B := make([]int, t)
	for i := 0; i < len(a); i++ {
		A[i] = a[i]
	}
	for i := 0; i < len(b); i++ {
		B[i] = b[i]
	}
	ntt(A, false)
	ntt(B, false)
	for i := 0; i < t; i++ {
		A[i] = mul(A[i], B[i])
	}
	ntt(A, true)
	Resize(&A, size)
	return A
}

func Resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

func ntt(a []int, rev bool) {
	size := len(a)
	if size == 1 {
		return
	}
	r := (MOD - 1) / size
	if rev {
		r = MOD - 1 - (MOD-1)/size
	}
	s := POWMOD(ROOT, r)
	kp := make([]int, size/2+1)
	for i := range kp {
		kp[i] = 1
	}
	for i := 0; i < size/2; i++ {
		kp[i+1] = mul(kp[i], s)
	}
	b := make([]int, size)
	for i, l := 1, size/2; i < size; i, l = i<<1, l>>1 {
		for j, r := 0, 0; j < l; j, r = j+1, r+i {
			for k, s := 0, kp[i*j]; k < i; k++ {
				p := a[k+r]
				q := a[k+r+size/2]
				b[k+2*r] = (p + q) % MOD
				b[k+2*r+i] = mul((p-q+MOD)%MOD, s)
			}
		}
		copy(a, b)
	}
	if rev {
		s = inverse(size)
		for i := 0; i < size; i++ {
			a[i] = mul(a[i], s)
		}
	}
}

func inverse(x int) int {
	return POWMOD(x, MOD-2)
}

func POWMOD(a, n int) int {
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
