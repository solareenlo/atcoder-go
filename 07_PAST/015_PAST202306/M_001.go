package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dp *big.Int
	dp = big.NewInt(0)

	var N int
	fmt.Fscan(in, &N)

	dp.SetBit(dp, 0, 1)
	sum := 0
	for i := 0; i < N-1; i++ {
		var d uint
		fmt.Fscan(in, &d)
		sum += int(d)
		dp = or(dp, lsh(dp, d))
	}

	ans := int(1e9)
	for i := max(0, sum/2-5000); i < sum/2+5000; i++ {
		if dp.Bit(i) != 0 {
			ans = min(ans, abs(sum-i*2))
		}
	}
	fmt.Println(ans)
}

const MX = 1000000

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
