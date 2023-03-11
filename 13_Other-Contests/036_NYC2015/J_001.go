package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	comb := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		comb[i] = make([]int, i+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			tmp := 0
			if j < i {
				tmp = comb[i-1][j]
			}
			comb[i][j] = tmp + comb[i-1][j-1]
			if comb[i][j] >= M {
				comb[i][j] -= M
			}
		}
	}

	ways := make([]int, N+1)
	for i := 0; i <= N; i += 2 {
		for j := 0; j <= i; j += 2 {
			sub := comb[i][j] * comb[j][j/2] % M * comb[i-j][(i-j)/2] % M
			ways[i] += sub
			if ways[i] >= M {
				ways[i] -= M
			}
		}
	}

	dp := make([]int, N+1)
	dp[0] = 1
	for i := 2; i <= N; i += 2 {
		for j := 2; j <= i; j += 2 {
			dp[i] += M - ways[j]*dp[i-j]%M
			if dp[i] >= M {
				dp[i] -= M
			}
		}
	}

	table := make([]int, N+1)
	for i := N; i >= 0; i-- {
		pw := 1
		for j := i; j <= N; j++ {
			table[j] = (table[j] + pw*dp[i]) % M
			pw = pw * 4 % M
		}
	}

	fpw := 1
	ans := 0
	for i := N; i >= 0; i-- {
		ans = (ans + fpw*table[i]) % M
		fpw = fpw * 4 % M
	}
	fmt.Println(ans)
}
