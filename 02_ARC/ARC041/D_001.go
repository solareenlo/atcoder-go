package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	G     = make([][2000][]int, 2)
	n     int
	m     int
	color = [2000]int{}
	cnt   int
)

func dfs(x, col int) bool {
	if ^color[x] != 0 {
		return color[x] != col
	}
	color[x] = col
	for _, y := range G[col][x] {
		cnt++
		tmp := 0
		if col == 0 {
			tmp = 1
		}
		if dfs(y, tmp) {
			return true
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)

	for j := 0; j < m; j++ {
		var a, b int
		var c string
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		tmp := 0
		if c == "b" {
			tmp = 1
		}
		G[tmp][a] = append(G[tmp][a], b)
		G[tmp][b] = append(G[tmp][b], a)
	}

	ok := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			for k := range color {
				color[k] = -1
			}
			cnt = 0
			if dfs(i, j) {
				ok = 1
			}
			tmp := 0
			if cnt == m {
				tmp = 1
			}
			ok |= tmp
		}
	}

	if ok != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
