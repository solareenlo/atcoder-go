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

	const N = 100005

	var n, k int
	fmt.Fscan(in, &n, &k)
	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l, r, ans := 0, k, 0
	for l <= r {
		mid := (l + r) / 2
		s := 0
		for i := 1; i <= n; i++ {
			s += min(a[i], mid)
		}
		if s > k {
			r = mid - 1
		} else {
			l = mid + 1
			ans = mid
		}
	}
	for i := 1; i <= n; i++ {
		t := min(a[i], ans)
		a[i] -= t
		k -= t
	}
	for i := 1; k > 0; i++ {
		if a[i] != 0 {
			a[i]--
			k--
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, a[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
