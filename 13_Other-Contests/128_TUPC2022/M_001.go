package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var K int
	fmt.Fscan(in, &K)
	N := make([]int, K)
	p := make([][]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &N[i])
		p[i] = make([]int, N[i])
		for j := range p[i] {
			p[i][j] = -1
		}
		for j := 1; j < N[i]; j++ {
			fmt.Fscan(in, &p[i][j])
			p[i][j]--
		}
	}
	M := 0
	for i := 0; i < K; i++ {
		M = max(M, N[i])
	}
	mt := rand.New(rand.NewSource(time.Now().UnixNano()))
	H := make([]int, M)
	for i := 0; i < M; i++ {
		H[i] = int(modulo(uint64(mt.Int63())))
	}
	hash := make([]int, K)
	for i := 0; i < K; i++ {
		c := make([][]int, N[i])
		for j := 1; j < N[i]; j++ {
			c[p[i][j]] = append(c[p[i][j]], j)
		}
		dp := make([]int, N[i])
		for j := range dp {
			dp[j] = 1
		}
		for j := 0; j < N[i]; j++ {
			for _, k := range c[j] {
				dp[k] = int(mul(uint64(dp[j]), uint64(H[len(c[j])])))
			}
		}
		a, b := 0, 0
		for j := 0; j < N[i]; j++ {
			if len(c[j]) != 0 {
				a = int(modulo(uint64(a + dp[j])))
			} else {
				b = int(modulo(uint64(b + dp[j])))
			}
		}
		hash[i] = int(mul(uint64(a), inverse(modulo(uint64(MOD+1-b)))))
	}
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			if hash[i] == hash[j] {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

const M30 = (1 << 30) - 1
const M31 = (1 << 31) - 1
const MOD = (1 << 61) - 1

func modulo(x uint64) uint64 {
	xu := x >> 61
	xd := x & MOD
	res := xu + xd
	if res >= MOD {
		res -= MOD
	}
	return res
}

func mul(a, b uint64) uint64 {
	au := a >> 31
	ad := a & M31
	bu := b >> 31
	bd := b & M31
	mid := au*bd + ad*bu
	midu := mid >> 30
	midd := mid & M30
	return modulo(au*bu*2 + midu + (midd << 31) + ad*bd)
}

func modpow(a, b uint64) uint64 {
	ans := uint64(1)
	for b > 0 {
		if b%2 == 1 {
			ans = mul(ans, a)
		}
		a = mul(a, a)
		b /= 2
	}
	return ans
}

func inverse(a uint64) uint64 {
	return modpow(a, MOD-2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
