package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k, m, sm, mx int
var a [101000]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &k, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sm += a[i]
		mx = max(mx, a[i])
	}
	if check() != 0 {
		fmt.Fprintln(out, "Chinatsu")
	} else {
		fmt.Fprintln(out, "Akari")
	}
	for i := 0; i < m-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		a[u] += v
		sm += v
		mx = max(mx, a[u])
		if check() != 0 {
			fmt.Fprintln(out, "Chinatsu")
		} else {
			fmt.Fprintln(out, "Akari")
		}
	}
}

func check() int {
	if n == 1 {
		if sm <= k {
			return 1
		} else {
			return (sm - k + 1) % 2
		}
	}
	if mx <= k || k%2 == 1 || mx <= sm-mx+k-1 {
		return sm % 2
	}
	return (sm + 1) % 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
