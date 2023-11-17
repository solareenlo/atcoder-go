package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	x := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &x[i])
	}

	sqrtN := int(math.Sqrt(float64(N)))

	val := make([]int, 0)
	for l, r := 1, 0; l <= N; l = r + 1 {
		r = N / (N / l)
		val = append(val, N/l)
	}

	K := len(val)
	dp := make([]int, K)
	dp[0] = 1
	for i := 0; i < M; i++ {
		if x[i] == 1 {
			continue
		}
		g := make([]int, K)
		for j := 0; val[j] >= x[i]; j++ {
			v := dp[j]*(1-x[i]) + g[j]
			nv := val[j] / x[i]
			if nv <= sqrtN {
				g[K-nv] += v
			} else {
				g[N/nv-1] += v
			}
		}
		for j := 0; j < K; j++ {
			dp[j] += g[j]
		}
	}

	ans := 0
	for i := 0; i < K; i++ {
		ans += val[i] * (val[i] + 1) / 2 * dp[i]
	}
	fmt.Println(ans)
}
