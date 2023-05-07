package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type edge struct {
		to    int
		color string
	}

	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		G := make([][]edge, n)
		for i := 0; i < n-1; i++ {
			var a, b int
			var c string
			fmt.Fscan(in, &a, &b, &c)
			a--
			b--
			G[a] = append(G[a], edge{b, c})
			G[b] = append(G[b], edge{a, c})
		}
		seen := make([]bool, n)
		var dfs func(int) float64
		dfs = func(node int) float64 {
			res := 0.0
			for _, e := range G[node] {
				if seen[e.to] {
					continue
				}
				seen[e.to] = true
				x := dfs(e.to)
				if e.color == "B" {
					k := max(1, int(math.Ceil(1-x)))
					res += (x + float64(k)) / math.Pow(2, float64(k-1))
				} else {
					k := max(1, int(math.Ceil(1+x)))
					res += (x - float64(k)) / math.Pow(2, float64(k)-1)
				}
			}
			return res
		}
		seen[0] = true
		val := dfs(0)
		if val > 0.0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
