package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n+1)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum[i] = sum[i-1] + a[i]
	}

	ans := 1 << 60
	for i := 1; i <= n; i++ {
		t := i*k + 5*(sum[n]-sum[n-i])
		for j := 2; i*(j-1) < n && t < ans; j++ {
			t += (j*2 + 1) * (sum[n-(j-1)*i] - sum[max(0, n-i*j)])
		}
		ans = min(ans, t)
	}

	fmt.Println(ans + n*k)
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
