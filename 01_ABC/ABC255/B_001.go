package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	const N = 1010
	a := make([]int, N)
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &a[i])
	}
	x := make([]float64, N)
	y := make([]float64, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	ans := 0.0
	for i := 1; i <= n; i++ {
		mx := 1e6 + 1
		for j := 1; j <= k; j++ {
			mx = min(mx, math.Sqrt((x[i]-x[a[j]])*(x[i]-x[a[j]])+(y[i]-y[a[j]])*(y[i]-y[a[j]])))
		}
		ans = max(ans, mx)
	}
	fmt.Println(ans)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
