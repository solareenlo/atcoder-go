package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K, A int
	fmt.Fscan(in, &N, &K, &A)

	var a, b, c, d [1500]int
	kei := A
	maxi := A
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i], &d[i])
		kei += a[i]
		maxi = max(maxi, a[i]+K/c[i]*(b[i]))
	}

	L := 1.0
	R := float64(maxi)
	M := float64(L+R) / 2.0
	seido := 1
	haba := (R - L) * 10000

	for float64(seido)*M <= haba {
		seido <<= 1
		sikii := M*float64(N+1) - float64(kei)
		if sikii <= 0 {
			L = M
			M = (L + R) / 2
			continue
		}

		var dp [2][1501]float64
		mae := dp[0]
		next := dp[1]
		for i := 0; i < N; i++ {
			for j := 0; j < K-c[i]+1; j++ {
				next[j+c[i]] = math.Max(next[j+c[i]], next[j]+float64(b[i]))
			}
			for j := K - d[i]; j >= 0; j-- {
				next[j+d[i]] = math.Max(next[j+d[i]], mae[j]+M-float64(a[i]))
			}
			for j := min(c[i], d[i]); j <= K; j++ {
				mae[j] = next[j]
			}
			mae, next = next, mae
			if sikii <= mae[K] {
				break
			}
		}

		if sikii <= mae[K] {
			L = M
		} else {
			R = M
		}
		M = (L + R) / 2
	}

	fmt.Println(M)
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
