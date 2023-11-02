package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dp [MX]*big.Int
	for i := range dp {
		dp[i] = big.NewInt(0)
	}

	var s, t, ans [10010]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var t1, t2 int
		fmt.Fscan(in, &t1, &t2)
		dp[t1].SetBit(dp[t1], t2, 1)
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &s[i], &t[i])
		ans[i] = -1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if dp[j].Bit(i) == 1 {
				dp[j] = or(dp[j], dp[i])
			}
		}
		for j := 1; j <= q; j++ {
			if ans[j] == -1 && dp[s[j]].Bit(t[j]) == 1 {
				ans[j] = max(s[j], t[j], i)
			}
		}
	}
	for i := 1; i <= q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

const MX = 2020

func lsh(A *big.Int, x uint) *big.Int {
	return new(big.Int).Lsh(A, x)
}

func rsh(A *big.Int, x uint) *big.Int {
	return new(big.Int).Rsh(A, x)
}

func and(A, B *big.Int) *big.Int {
	return new(big.Int).And(A, B)
}

func or(A, B *big.Int) *big.Int {
	return new(big.Int).Or(A, B)
}

func xor(A, B *big.Int) *big.Int {
	return new(big.Int).Xor(A, B)
}

func FindFirst(b *big.Int) int {
	for i := 0; i < MX; i++ {
		if b.Bit(i) == 1 {
			return i
		}
	}
	return MX
}

func FindNext(i int, b *big.Int) int {
	for j := i + 1; j < MX; j++ {
		if b.Bit(j) == 1 {
			return j
		}
	}
	return MX
}

func BitCount(n *big.Int) int {
	count := 0
	for _, v := range n.Bits() {
		count += popcount(uint64(v))
	}
	return count
}

func popcount(x uint64) int {
	const (
		m1  = 0x5555555555555555
		m2  = 0x3333333333333333
		m4  = 0x0f0f0f0f0f0f0f0f
		h01 = 0x0101010101010101
	)
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return int((x * h01) >> 56)
}
