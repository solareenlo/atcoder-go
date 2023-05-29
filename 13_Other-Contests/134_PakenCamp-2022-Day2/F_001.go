package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var N int
	fmt.Fscan(in, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
		if P[i] != -1 {
			P[i]--
		}
	}
	if N == 1 {
		fmt.Println(1)
		return
	}
	if N == 2 {
		if P[0] == 0 || P[1] == 1 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
		return
	}
	for i := 0; i < N; i++ {
		if P[i] == i {
			fmt.Println(0)
			return
		}
	}
	C := 0
	for i := range P {
		if P[i] == -1 {
			C++
		}
	}
	ans := 1
	for i := 0; i < C; i++ {
		ans = ans * (N - 1) % MOD
	}
	used := make([]bool, N)
	for i := 0; i < N; i++ {
		if P[i] != -1 {
			if used[P[i]] {
				fmt.Println(ans)
				return
			}
			used[P[i]] = true
		}
	}
	A := 0
	for i := 0; i < N; i++ {
		if P[i] == -1 && !used[i] {
			A++
		}
	}
	B := C - A
	dp := make([]int, A+1)
	dp[0] = 1
	for i := 0; i < B; i++ {
		dp[0] = dp[0] * (i + 1) % MOD
	}
	for i := 0; i < A; i++ {
		dp[i+1] = dp[i] * (B + i) % MOD
		if i > 0 {
			dp[i+1] = (dp[i+1] + dp[i-1]*i%MOD) % MOD
		}
	}
	fmt.Println((ans - dp[A] + MOD) % MOD)
}
