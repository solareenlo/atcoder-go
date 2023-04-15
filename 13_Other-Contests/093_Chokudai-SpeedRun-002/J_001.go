package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	G1 := A[0]
	G2 := B[0]
	for i := 1; i < N; i++ {
		G1 = max(gcd(G1, A[i]), gcd(G1, B[i]))
		G2 = max(gcd(G2, A[i]), gcd(G2, B[i]))
	}
	fmt.Println(max(G1, G2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
