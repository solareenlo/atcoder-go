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

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 0; i < n; i++ {
		fmt.Print(i+1, ": {")
		for j := 0; j < len(g[i]); j++ {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(g[i][j] + 1)
		}
		fmt.Println("}")
	}
}
