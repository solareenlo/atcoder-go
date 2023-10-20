package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = math.MaxInt - 1

	var n, m, c int
	fmt.Fscan(in, &n, &m, &c)
	a := make([]int, n)
	b := make([]int, n)
	a_max := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		a_max = max(a_max, a[i])
	}
	ans := INF
	dp := make([]int, pow(a_max+3, 2))
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for w := 0; w < a_max; w++ {
		step := INF
		best_a := 1
		best_b := INF
		for i := 0; i < n; i++ {
			if w+1 <= a[i] {
				step = min(step, b[i])
			}
			cand_a := min(w+1, a[i])
			cand_b := b[i]
			if float64(cand_b)/float64(cand_a) < float64(best_b)/float64(best_a) {
				best_a = cand_a
				best_b = cand_b
			}
		}
		for t := 0; t < len(dp); t++ {
			if len(dp) <= t+w+1 {
				break
			}
			dp[t+w+1] = min(dp[t+w+1], dp[t]+step)
		}
		k := max(0, (m-len(dp))/best_a+1)
		ans = min(ans, c*w+k*best_b+dp[m-k*best_a])
	}
	fmt.Println(ans)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
