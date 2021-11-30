package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, c int
	fmt.Fscan(in, &n, &m, &c)

	a := [1002][1002]int{}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for i := 1; i < m+1; i++ {
		a[0][i] = 1 << 60
	}

	res := 1 << 60
	for i := 1; i < n+1; i++ {
		s, mini := 0, 1<<60
		for j := 1; j < m+1; j++ {
			res = min(res, a[i][j]+a[i-1][j]+c)
			a[i][j] = min(a[i][j], a[i-1][j]+c)
			res = min(res, mini+s+a[i][j])
			mini = min(mini, a[i][j]-s)
			s += c
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
