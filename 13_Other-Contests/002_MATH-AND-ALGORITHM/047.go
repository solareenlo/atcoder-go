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

	p := make([]int, n)
	for i := range p {
		p[i] = -1
	}
	que := make([]int, 0)
	for i := 0; i < n; i++ {
		if p[i] != -1 {
			continue
		}
		p[i] = 0
		que = append(que, i)
		for len(que) > 0 {
			a := que[0]
			que = que[1:]
			for _, na := range g[a] {
				if p[na] == p[a] {
					fmt.Println("No")
					return
				}
				if p[na] == -1 {
					p[na] = p[a] ^ 1
					que = append(que, na)
				}
			}
		}
	}
	fmt.Println("Yes")
}
