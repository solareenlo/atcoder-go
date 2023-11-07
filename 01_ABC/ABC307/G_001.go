package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const N = 5050

	var a, s [N]int
	var f [N][N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s[i] = a[i] + s[i-1]
	}
	k := s[n] / n
	for k*n > s[n] {
		k--
	}
	r := s[n] - k*n
	for i := range f {
		for j := range f[i] {
			f[i][j] = INF
		}
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= r; j++ {
			f[i][j] = f[i-1][j] + abs(k*i+j-s[i])
			if j != 0 {
				f[i][j] = min(f[i][j], f[i-1][j-1]+abs(k*i+j-s[i]))
			}
		}
	}
	fmt.Println(f[n][r])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
