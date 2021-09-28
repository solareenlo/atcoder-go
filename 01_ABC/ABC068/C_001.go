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

	c := make([][]int, n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		a--
		b--
		c[a] = append(c[a], b)
	}

	for i := 0; i < len(c[0]); i++ {
		for j := 0; j < len(c[c[0][i]]); j++ {
			if c[c[0][i]][j] == n-1 {
				fmt.Println("POSSIBLE")
				return
			}
		}
	}
	fmt.Println("IMPOSSIBLE")
}
