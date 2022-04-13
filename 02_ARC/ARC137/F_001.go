package main

import "fmt"

const N = 1050000

var (
	R   = make([]int, N)
	h   = make([]int, N)
	G   = make([]int, N)
	k   int
	lim int
	A   = make([]int, N)
	B   = make([]int, N)
	W   = make([]int, N)
	f   = make([]int, N)
)

func Dft(a []int, lim, o int) {
	for i := 0; i < lim; i++ {
		if R[i] > i {
			a[i], a[R[i]] = a[R[i]], a[i]
		}
	}
	for i := 1; i < lim; i <<= 1 {
		W[0] = 1
		W[1] = powMod(3, (mod-1)/i/2)
		for j := 2; j < i; j++ {
			W[j] = W[j-1] * W[1] % mod
		}
		for j := 0; j < lim; j += (i << 1) {
			for k := 0; k < i; k++ {
				t := a[i+j+k] * W[k] % mod
				a[i+j+k] = (a[j+k] - t + mod) % mod
				a[j+k] += t
				a[j+k] %= mod
			}
		}
	}
	if o != 0 {
		tmp := a[1:lim]
		tmp = reverseOrderInt(tmp)
		for i := 0; i < lim-1; i++ {
			a[i+1] = tmp[i]
		}
		t := powMod(lim, mod-2)
		for i := 0; i < lim; i++ {
			a[i] *= t
			a[i] %= mod
		}
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func fen(l, r int) {
	if l == r {
		if l < k {
			if l&1 != 0 {
				h[l] = 0
			} else {
				h[l] = G[l/2]
			}
		}
		return
	}
	mid := (l + r) >> 1
	len1 := mid - l
	len2 := r - l

	fen(l, mid)
	if r < k {
		fen(mid+1, r)
		return
	}

	for lim = 1; lim <= r-l; lim <<= 1 {

	}
	for i := 1; i < lim; i++ {
		R[i] = (R[i>>1] >> 1) | ((lim >> 1) * (i & 1))
	}
	for i := 0; i < len1+1; i++ {
		A[i] = h[i+l]
	}
	for i := len1 + 1; i < lim-len1; i++ {
		A[i] = 0
	}
	for i := 0; i < len2+1; i++ {
		B[i] = f[i]
	}
	for i := len2 + 1; i < lim-len2; i++ {
		B[i] = 0
	}
	Dft(A, lim, 0)
	Dft(B, lim, 0)
	for i := 0; i < lim; i++ {
		A[i] *= B[i]
		A[i] %= mod
	}
	Dft(A, lim, 1)

	for i := mid + 1; i < r+1; i++ {
		h[i] += A[i-l]
		h[i] %= mod
	}
	fen(mid+1, r)
}

func main() {
	jc := make([]int, N)
	inv := make([]int, N)
	jc[0] = 1
	jc[1] = 1
	inv[0] = 1
	inv[1] = 1
	G[0] = 1
	G[1] = 1
	for i := 2; i < 200011; i++ {
		jc[i] = jc[i-1] * i % mod
		inv[i] = inv[mod%i] * (mod - mod/i) % mod
		G[i] = G[i-1] * (2*i - 1) % mod
	}
	for i := 2; i < 200011; i++ {
		inv[i] *= inv[i-1]
		inv[i] %= mod
	}

	var n int
	fmt.Scan(&n, &k)
	k++

	for j := 0; j < k/2+1; j++ {
		var tmp int
		if j&1 != 0 {
			tmp = -1
		} else {
			tmp = 1
		}
		f[k-2*j] = tmp * jc[k] * inv[2*j] % mod * inv[k-2*j] % mod * G[j] % mod
	}
	for i := 0; i < k; i++ {
		f[i] = (mod - f[i]) % mod
	}
	for i := 0; i < k/2+1; i++ {
		f[i], f[k-i] = f[k-i], f[i]
	}
	f[0] = 0

	fen(0, 2*n)
	ans := h[2*n]
	ans = (ans%mod + mod) % mod
	fmt.Println(ans * powMod(G[n], mod-2) % mod)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
