package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [50][50]int
	for i := 0; i < m; i++ {
		var t, b int
		fmt.Fscan(in, &t, &b)
		t--
		b--
		a[t][b] = 1
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a[i][j] |= a[i][k] & a[k][j]
			}
		}
	}

	cnt := 1
	for u := 0; u < n; u++ {
		if a[u][0] != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
