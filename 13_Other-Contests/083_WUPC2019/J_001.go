package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var N, M int
	fmt.Fscan(in, &N, &M)

	fact := make([]int, M+1)
	rev := make([]int, M+1)
	fact_rev := make([]int, M+1)
	for i := 0; i < M+1; i++ {
		fact[i] = 1
		rev[i] = 1
		fact_rev[i] = 1
	}
	for i := 2; i < M+1; i++ {
		fact[i] = (fact[i-1] * i) % mod
		rev[i] = (mod - ((mod/i)*rev[mod%i])%mod)
		fact_rev[i] = (fact_rev[i-1] * rev[i]) % mod
	}
	var comb func(int, int) int
	comb = func(a, b int) int {
		return (fact[a] * ((fact_rev[b] * fact_rev[a-b]) % mod)) % mod
	}
	var mul func([]int, []int) []int
	mul = func(l, r []int) []int {
		a := len(l)
		b := len(r)
		p := make([]int, a+b-1)
		for i := 0; i < a; i++ {
			for j := 0; j < b; j++ {
				p[i+j] = (p[i+j] + l[i]*r[j]) % mod
			}
		}
		return p
	}

	dp := make([]int, 1)
	dp[0] = 1
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		tmp := make([]int, a+1)
		for x := 0; x < a+1; x++ {
			tmp[x] = comb(a, x)
			tmp[x] = (tmp[x] * fact_rev[x]) % mod
		}
		dp = mul(dp, tmp)
	}
	ans := 0
	for i := 0; i < M+1; i++ {
		ans = (mod - ans)
		ans = (ans + fact[i]*dp[i]) % mod
	}
	fmt.Println(ans)
}
