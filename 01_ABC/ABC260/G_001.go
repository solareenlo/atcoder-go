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

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 2000
	g := make([]string, N)
	var ans [N][N]int
	var pre1, pre2 [8*N + 1]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &g[i])
		for j := 0; j < n; j++ {
			if g[i][j] == 'O' {
				pre1[j]++
				pre2[j+2*m+2*i]--
			}
			if i >= m && g[i-m][j] == 'O' {
				pre1[j]--
				pre2[j+2*m+2*(i-m)]++
			}
		}
		cur := 0
		for j := 0; j < n; j++ {
			cur += pre1[j] + pre2[j+2*i]
			ans[i][j] = cur
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		var i, j int
		fmt.Fscan(in, &i, &j)
		fmt.Fprintln(out, ans[i-1][j-1])
		q--
	}
}
