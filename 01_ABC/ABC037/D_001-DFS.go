package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	h, w  int
	a, dp [][]int
	dx    [4]int = [4]int{0, 1, 0, -1}
	dy    [4]int = [4]int{1, 0, -1, 0}
	mod   int    = int(1e9 + 7)
)

func dfs(y, x int) int {
	if dp[y][x] > 0 {
		return dp[y][x]
	}
	res := 1
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if !(nx < 0 || ny < 0 || nx >= w || ny >= h || a[ny][nx] <= a[y][x]) {
			res = (res + dfs(ny, nx)) % mod
		}
	}
	dp[y][x] = res
	return dp[y][x]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &h, &w)
	a = make([][]int, h)
	dp = make([][]int, h)
	for y := 0; y < h; y++ {
		a[y] = make([]int, w)
		dp[y] = make([]int, w)
		for x := 0; x < w; x++ {
			fmt.Fscan(in, &a[y][x])
		}
	}

	res := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			res = (res + dfs(y, x)) % mod
		}
	}
	fmt.Fprintln(out, res)
}
