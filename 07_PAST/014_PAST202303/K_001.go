package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, t int
	var p float64
	fmt.Fscan(in, &N, &t, &p)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := make([]float64, N+2)

	for i := N - 1; i >= 0; i-- {
		coins := (a[i] * t) / 100
		dp[i] = math.Max(dp[i], dp[i+1]+float64(a[i]-coins))
		pp := p
		pp /= 100.0
		dp[i] = math.Max(dp[i], (1.0-pp)*(dp[i+1]+float64(a[i]))+pp*(dp[i+2]+float64(a[i]-coins)))
	}

	fmt.Printf("%.20f\n", dp[0])
}
