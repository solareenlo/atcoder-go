package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	k = min(k, 2*n+k%2)

	var s string
	fmt.Fscan(in, &s)
	a := make([]int, 200002)
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			a[i] = 1
		}
	}

	now := 0
	tag := 0
	for i := 1; i <= k; i++ {
		if a[now]^tag != 0 {
			a[now] ^= 1
		} else {
			now = (now + 1) % n
			tag ^= 1
		}
	}

	for i, j := now, 0; j < n; j++ {
		fmt.Fprint(out, string("BA"[a[i]^tag]))
		i = (i + 1) % n
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
