package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200010

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a, b [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = b[i-1] + i*a[i]
		a[i] += a[i-1]
	}

	ans := -1_000_000_000_000_000_000
	for i := m; i <= n; i++ {
		ans = max(ans, b[i]-b[i-m]-(a[i]-a[i-m])*(i-m))
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
