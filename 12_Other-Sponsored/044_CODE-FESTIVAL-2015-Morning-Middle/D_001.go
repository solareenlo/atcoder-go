package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100000
	var a, st, c [N + 5]int
	var top int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	top++
	st[top] = 0
	for i := 1; i <= n+1; i++ {
		for a[st[top]] > a[i] {
			c[i-st[top-1]-1] += a[st[top]] - max(a[st[top-1]], a[i])
			top--
		}
		top++
		st[top] = i
	}
	ans := n - 2*k
	for i := 1; i <= n && k > 0; i++ {
		if i*c[i] <= k {
			k -= i * c[i]
			ans -= c[i] * 2
		} else {
			ans -= k / i * 2
			k = 0
		}
	}
	for i := 1; i <= n+1; i++ {
		ans += 2 * a[i]
		if a[i] > a[i-1] {
			ans += 2 * (a[i] - a[i-1])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
