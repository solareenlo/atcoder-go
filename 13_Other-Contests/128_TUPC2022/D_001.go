package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, P int
	fmt.Fscan(in, &N, &P)
	inv := make([]int, N+1)
	inv[1] = 1
	for i := 2; i <= N; i++ {
		inv[i] = P - inv[P%i]*(P/i)%P
	}
	fact := make([]int, N+1)
	finv := make([]int, N+1)
	fact[0] = 1
	finv[0] = 1
	for i := 1; i <= N; i++ {
		fact[i] = fact[i-1] * i % P
		finv[i] = finv[i-1] * inv[i] % P
	}
	sum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j += i {
			sum[j] += i
			sum[j] %= P
		}
	}
	ans := 0
	for i := 0; i < N+1; i++ {
		ans += sum[i] * sum[N-i] % P * fact[N] % P * finv[i] % P * finv[N-i] % P
	}
	ans %= P
	fmt.Println(ans)
}
