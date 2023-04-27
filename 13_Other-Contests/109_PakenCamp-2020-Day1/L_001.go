package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var Q, X int
	fmt.Fscan(in, &Q, &X)

	A := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &A[i])
	}

	pow2 := 1
	sum := X
	for i := 0; i < Q; i++ {
		m := A[i]
		p := 1
		for m > 0 {
			p = p * 10 % MOD
			m /= 10
		}
		a := (sum + pow2*A[i]%MOD) % MOD
		b := (sum*p%MOD + pow2*A[i]%MOD) % MOD
		sum = (a + b) % MOD
		pow2 = pow2 * 2 % MOD
	}
	fmt.Println(sum)
}
