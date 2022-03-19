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
	s := make([]int, n+1)
	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s[i] = s[i-1] + a[i]
		if a[i] > 0 {
			t[i] = t[i-1] + a[i]
		} else {
			t[i] = t[i-1]
		}
	}

	ans := 0
	for i := 1; i <= n-k+1; i++ {
		ans = max(ans, max(s[i+k-1]-s[i-1], 0)+t[n]-t[i+k-1]+t[i-1])
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
