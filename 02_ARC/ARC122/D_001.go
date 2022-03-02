package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 12000005
const K = 30

var (
	g     = [N][2]int{}
	t int = 1
	s     = [N]int{}
)

func I(v int) {
	x := 1
	for i := K; i >= 0; i-- {
		j := v >> i & 1
		if g[x][j] == 0 {
			t++
			g[x][j] = t
		}
		x = g[x][j]
		s[x]++
	}
}

func Q(x, y, d int) int {
	if d < 0 {
		return 0
	}
	mini := 1 << 60
	for j := 0; j < 2; j++ {
		if g[x][j] != 0 {
			if g[y][j] != 0 {
				mini = min(mini, Q(g[x][j], g[y][j], d-1))
			} else {
				mini = min(mini, Q(g[x][j], g[y][j^1], d-1)^(1<<d))
			}
		}
	}
	return mini
}

func W(x, d int) int {
	if x == 0 || d < 0 {
		return 0
	}
	if s[g[x][0]]&1 != 0 {
		return Q(g[x][0], g[x][1], d-1) ^ (1 << d)
	}
	return max(W(g[x][0], d-1), W(g[x][1], d-1))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	n *= 2

	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		I(x)
	}
	fmt.Println(W(1, K))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
