package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)

	g := make([][]int, n+1)
	c := make([]string, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := n - 1; i > 0; i-- {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var f func(int, int) [2]int
	f = func(u, v int) [2]int {
		X, Y := 1, 1
		for _, w := range g[v] {
			if w != u {
				tmp := f(v, w)
				x := tmp[0]
				y := tmp[1]
				X = X * ((x + x + y) % MOD) % MOD
				if c[v] == c[w] {
					Y = Y * ((x + y) % MOD) % MOD
				} else {
					Y = Y * x % MOD
				}
			}
		}
		return [2]int{(X - Y + MOD) % MOD, Y}
	}
	fmt.Println(f(0, 1)[0])
}
