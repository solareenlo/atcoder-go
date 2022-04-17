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
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	const mod = 1_000_000_007
	ans := 1
	for i := 1; i <= n; i++ {
		if a[i] > a[i-1] && b[i] > b[i+1] {
			if a[i] != b[i] {
				ans = 0
			}
		} else if a[i] > a[i-1] {
			if a[i] > b[i] {
				ans = 0
			}
		} else if b[i] > b[i+1] {
			if b[i] > a[i] {
				ans = 0
			}
		} else {
			ans = ans * min(a[i], b[i]) % mod
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
