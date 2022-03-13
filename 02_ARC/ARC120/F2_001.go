package main

import (
	"bufio"
	"fmt"
	"os"
)

func upd(x int) int    { return x + ((x >> 31) & mod) }
func Add(x, y int) int { return upd(x + y - mod) }

const mod = 998244353
const X = 1<<20 | 10
const size = X

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}

var (
	imaxw = [30]int{}
	Wn    = make([][]int, 32)
)

func prepare(x int) {
	for i := 0; i <= 25; i++ {
		imaxw[i] = powMod(2, mod-1-i)
	}
	var MXW int
	for MXW = 1; MXW <= x; MXW <<= 1 {
	}
	for i := 1; i < 27; i++ {
		if (1 << i) > MXW {
			break
		}
		Wn[i] = make([]int, (1<<i)+1)
		Wn[i][0] = 1
		v := powMod(3, (mod-1)>>i)
		for j := 1; j < (1<<i)+1; j++ {
			Wn[i][j] = Wn[i][j-1] * v % mod
		}
	}
}

var (
	maxw int
	btw  int
	I    int
	id   = make([]int, X)
)

func INIT(x, tp int) {
	for maxw, btw = 1, 0; maxw <= x; btw++ {
		maxw <<= 1
	}
	I = imaxw[btw]
	if tp != 0 {
		for i, s := 1, btw-1; i < maxw; i++ {
			id[i] = (id[i>>1] >> 1) | ((i & 1) << s)
		}
	}
}

func ntt(w []int, tp int) {
	for i := 0; i < maxw; i++ {
		if i < id[i] {
			w[i], w[id[i]] = w[id[i]], w[i]
		}
	}
	for md, len, b := 1, 2, 1; md < maxw; b++ {
		for l := 0; l < maxw; l += len {
			for u := 0; u < md; u++ {
				x := w[u+l]
				tmp := len - u
				if ^tp != 0 {
					tmp = u
				}
				y := w[u+l+md] * Wn[b][tmp] % mod
				w[u+l] = upd(x + y - mod)
				w[u+l+md] = upd(x - y)
			}
		}
		md = len
		len <<= 1
	}
}

var (
	D  int
	AA = make([]int, X)
	BB = make([]int, X)
)

func solve(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	solve(l, mid)
	solve(mid+1, r)
	INIT((r-l+1)*(D-1), 1)
	A := make([]int, maxw<<2)
	B := make([]int, maxw<<2)
	C := make([]int, maxw<<2)
	DD := make([]int, maxw<<2)
	f := (l - 2) * D
	for i := 0; i < (mid-l+1)*(D-1)+1; i++ {
		A[i] = AA[f+i]
	}
	f = (mid + 1 - 2) * D
	for i := 0; i < (r-mid)*(D-1)+1; i++ {
		B[i] = AA[f+i]
	}
	f = (l - 2) * D
	for i := 0; i < (mid-l+1)*(D-1)+1; i++ {
		C[i] = BB[f+i]
	}
	f = (mid + 1 - 2) * D
	for i := 0; i < (r-mid)*(D-1)+1; i++ {
		DD[i] = BB[f+i]
	}
	ntt(A, 1)
	ntt(B, 1)
	ntt(C, 1)
	ntt(DD, 1)
	for i := 0; i < maxw; i++ {
		C[i] = (DD[i] + C[i]*B[i]) % mod
		A[i] = A[i] * B[i] % mod
	}
	ntt(A, -1)
	ntt(C, -1)
	f = (l - 2) * D
	for i := 0; i < (r-l+1)*(D-1)+1; i++ {
		AA[f+i] = A[i] * I % mod
	}
	f = (l - 2) * D
	for i := 0; i < (r-l+1)*(D-1)+1; i++ {
		BB[f+i] = C[i] * I % mod
	}
}

func getVal(n, k int) int {
	if k > (n+D-1)/D {
		os.Exit(1)
	}
	return nCrMod(n+(D-1)*(1-k), k)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K, &D)

	if K == 1 {
		z := 0
		for i := 1; i < N+1; i++ {
			var a int
			fmt.Fscan(in, &a)
			z = Add(z, a)
		}
		fmt.Println(z)
		return
	}

	w := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		fmt.Fscan(in, &w[i])
	}
	prepare(N + 10)
	initMod()
	tag := 1
	for i := 2; i < K+1; i++ {
		tmp := (i - 2) * D
		for j := 1; j < D; j++ {
			AA[tmp+j] = mod - 1
		}
		for j := 0; j < D-1; j++ {
			BB[tmp+j] = (D - 1 - j) * tag % mod
		}
		tag = tag * upd(1-D) % mod
		tag = Add(tag, getVal(N-(K-i+1)*D, i-1))
	}

	F := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		F[i] = tag
	}
	solve(2, K)
	for i := 0; i < (D-1)*(K-1)+1; i++ {
		F[i+1] = Add(F[i+1], BB[i])
		F[N-i] = Add(F[N-i], BB[i])
	}

	z := 0
	for i := 1; i < N+1; i++ {
		z = Add(z, w[i]*F[i]%mod)
	}
	fmt.Println(z)
}
