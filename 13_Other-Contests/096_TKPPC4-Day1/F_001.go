package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a, b [500][500]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &b[i][j])
		}
	}

	t := 0
	for i := 0; i < n; i++ {
		t2 := INF
		for j := 0; j < m; j++ {
			t2 = min(t2, (t+a[i][j]-1)/a[i][j]*a[i][j]+b[i][j])
		}
		t = t2
	}
	fmt.Println(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
