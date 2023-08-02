package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A [16]int
	var B, C [50]int
	var I [50][16]int

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &B[i], &C[i])
		for j := 0; j < C[i]; j++ {
			fmt.Fscan(in, &I[i][j])
			I[i][j]--
		}
	}
	ans := 0
	for i := 0; i < 1<<N; i++ {
		if popcount(uint32(i)) != 9 {
			continue
		}
		now := 0
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				now += A[j]
			}
		}
		for j := 0; j < M; j++ {
			cnt := 0
			for k := 0; k < C[j]; k++ {
				if ((i >> I[j][k]) & 1) != 0 {
					cnt++
				}
			}
			if cnt >= 3 {
				now += B[j]
			}
		}
		ans = max(ans, now)
	}
	fmt.Println(ans)
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
