package main

import (
	"bufio"
	"fmt"
	"os"
)

const S = 32

var M, N int
var a, Len [130]int

func contain(a, lena, b, lenb int) bool {
	for i := 0; i <= lena-lenb; i, a = i+1, a>>1 {
		if (a ^ (a >> lenb << lenb)) == b {
			return true
		}
	}
	return false
}

func check(s, lens int) bool {
	for i := 1; i <= M; i++ {
		if contain(s, lens, a[i], Len[i]) {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	K = S
	T := make([][]int, K)
	for i := 0; i < K; i++ {
		T[i] = make([]int, K)
	}
	var f, g [32]int

	fmt.Fscan(in, &N, &M)
	for i := 1; i <= M; i++ {
		var s string
		fmt.Fscan(in, &s)
		var j int
		for j = 0; j < len(s); j++ {
			tmp := 0
			if s[j] == 'b' {
				tmp = 1
			}
			a[i] = a[i]<<1 | tmp
		}
		Len[i] = j
	}
	if N <= 5 {
		ans := 0
		for s := 0; s < (1 << N); s++ {
			if check(s, N) {
				ans++
			}
		}
		fmt.Println(ans)
		return
	}
	for s := 0; s < 32; s++ {
		if check(s<<1, 6) {
			T[(s<<1)&31][s] = 1
		}
		if check(s<<1|1, 6) {
			T[(s<<1|1)&31][s] = 1
		}
		if check(s, 5) {
			f[s] = 1
		}
	}
	t := matPowMod(T, N-5)
	for i := 0; i < S; i++ {
		for j := 0; j < S; j++ {
			g[i] = (g[i] + t[i][j]*f[j]%MOD) % MOD
		}
	}
	ans := 0
	for i := 0; i < S; i++ {
		ans = (ans + g[i]) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353

var K int

func matMulMod(A, B [][]int) [][]int {
	C := make([][]int, K)
	for i := range C {
		C[i] = make([]int, K)
	}
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j] % MOD
				C[i][j] %= MOD
			}
		}
	}
	return C
}

func matPowMod(A [][]int, n int) [][]int {
	T := make([][]int, K)
	for i := range T {
		T[i] = make([]int, K)
	}
	if n == 0 {
		for i := 0; i < K; i++ {
			for j := 0; j < K; j++ {
				if i == j {
					T[i][j] = 1
				} else {
					T[i][j] = 0
				}
			}
		}
		return T
	}
	T = matPowMod(A, n>>1)
	T = matMulMod(T, T)
	if n&1 != 0 {
		T = matMulMod(T, A)
	}
	return T
}
