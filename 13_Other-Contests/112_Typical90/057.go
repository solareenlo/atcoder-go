package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var A [301]*big.Int
	for i := range A {
		A[i] = big.NewInt(0)
	}
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		for t > 0 {
			t--
			var x int
			fmt.Fscan(in, &x)
			x--
			A[x].SetBit(A[x], i, 1)
		}
	}

	for i := 0; i < m; i++ {
		var x uint
		fmt.Fscan(in, &x)
		A[i].SetBit(A[i], n, x)
	}

	rank := 0
	for i := 0; i < n; i++ {
		pivot := -1
		for j := rank; j < m; j++ {
			if A[j].Bit(i) != 0 {
				pivot = j
				break
			}
		}
		if pivot == -1 {
			continue
		}
		A[pivot], A[rank] = A[rank], A[pivot]
		for j := 0; j < m; j++ {
			if j != rank && A[j].Bit(i) != 0 {
				A[j].Xor(A[j], A[rank])
			}
		}
		rank++
	}
	for i := rank; i < m; i++ {
		if A[i].Bit(n) != 0 {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(powMod(2, n-rank))
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
