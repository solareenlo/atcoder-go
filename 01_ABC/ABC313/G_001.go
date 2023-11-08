package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var N int
	fmt.Fscan(in, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	sort.Ints(A)
	ans := 0
	B := 0
	ans = (ans + A[0] + 1) % MOD
	for i := 0; i < N-1; i++ {
		B += A[i]
		b := B + (N-i-1)*(A[i]+1)
		ans = (ans + (A[i+1]-A[i])*(b/N+1)%MOD) % MOD
		ans = (ans + floor_sum(A[i+1]-A[i], N, N-i-1, b%N)) % MOD
	}
	fmt.Println(ans)
}

func floor_sum(n, m, a, b int) int {
	ans := 0
	if a >= m {
		ans += (n - 1) * n * (a / m) / 2
		a %= m
	}
	if b >= m {
		ans += n * (b / m)
		b %= m
	}

	yMax := (a*n + b) / m
	xMax := (yMax*m - b)
	if yMax == 0 {
		return ans
	}
	ans += (n - (xMax+a-1)/a) * yMax
	ans += floor_sum(yMax, a, m, (a-xMax%a)%a)
	return ans
}
