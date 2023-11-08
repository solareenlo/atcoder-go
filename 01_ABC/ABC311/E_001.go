package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3005

	var a [N][N]bool
	var f [N][N]int

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[x][y] = true
	}
	ans := 0
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if !a[i][j] {
				f[i][j] = min(f[i-1][j-1], f[i][j-1], f[i-1][j]) + 1
				ans += f[i][j]
			}
		}
	}
	fmt.Println(ans)
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
