package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005

	var s, a [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i], &a[i])
	}
	for i := n; i > 0; i-- {
		a[s[i]] += max(a[i], 0)
	}
	mx := -N
	for i := 1; i <= n; i++ {
		mx = max(mx, a[i])
	}
	fmt.Println(mx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
