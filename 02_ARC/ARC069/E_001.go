package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	type pair struct{ q, p int }
	a := make([]pair, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].q)
		a[i].p = i
	}

	b := a[1 : n+1]
	sort.Slice(b, func(i, j int) bool {
		if b[i].q != b[j].q {
			return b[i].q > b[j].q
		}
		return b[i].p > b[j].p
	})

	ans := make([]int, 100002)
	x := 1 << 60
	for i := 1; i <= n; i++ {
		x = min(x, a[i].p)
		ans[x] += (a[i].q - a[i+1].q) * i
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
