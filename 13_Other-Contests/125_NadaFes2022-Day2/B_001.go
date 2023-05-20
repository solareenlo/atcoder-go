package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var a, b [200010]int

	var n int
	fmt.Fscan(in, &n)
	mx := -INF
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	ans := INF
	mn := INF
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &b[i])
		mn = min(mn, b[i])
		ans = min(ans, a[i]*b[i])
	}
	ans = min(ans, mx*mn)
	fmt.Println(ans)
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
