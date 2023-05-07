package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, S int
	fmt.Fscan(in, &N, &S)

	var A, B [100]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}

	var dp [101]*big.Int
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	dp[N].SetBit(dp[N], 0, 1)
	for i := N - 1; i >= 0; i-- {
		dp[i] = or(lsh(dp[i+1], uint(A[i])), lsh(dp[i+1], uint(B[i])))
	}
	if dp[0].Bit(S) != 0 {
		for i := 0; i < N; i++ {
			if S >= A[i] && dp[i+1].Bit(S-A[i]) != 0 {
				fmt.Print("A")
				S -= A[i]
			} else {
				fmt.Print("B")
				S -= B[i]
			}
		}
		fmt.Println()
	} else {
		fmt.Println("Impossible")
	}
}

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
