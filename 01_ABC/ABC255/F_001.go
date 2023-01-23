package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200010

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var c int
var a, b, pa, pb, lc, rc [N]int

func main() {
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		pa[a[i]] = i
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		pb[b[i]] = i
	}

	if dfs(1, n) != 1 {
		fmt.Println(-1)
		os.Exit(0)
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d %d\n", lc[i], rc[i])
	}
}
func dfs(l, r int) int {
	if l > r {
		return 0
	}
	c++
	m := pb[a[c]]
	x := a[c]
	if m < l || m > r {
		fmt.Println(-1)
		os.Exit(0)
	}
	lc[x] = dfs(l, m-1)
	rc[x] = dfs(m+1, r)
	return x
}
