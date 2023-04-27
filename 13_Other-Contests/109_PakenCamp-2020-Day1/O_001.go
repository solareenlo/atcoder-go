package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var a, b [500010]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	bas := make([]int, 0)
	for i := 0; i < n; i++ {
		c := a[i]*1024 + b[i]
		for j := 0; j < len(bas); j++ {
			c = min(c, c^bas[j])
		}
		if c != 0 {
			bas = append(bas, c)
		}
	}

	k := len(bas)
	ans := 0
	for i := 0; i < (1 << k); i++ {
		x := 0
		for j := 0; j < k; j++ {
			if ((i >> j) & 1) != 0 {
				x ^= bas[j]
			}
		}
		ans = max(ans, x%1024+x/1024)
	}
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
