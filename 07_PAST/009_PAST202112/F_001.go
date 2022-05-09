package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	s    = make([]string, 3)
	used = [10][10]bool{}
)

func dfs(x, y int) {
	if used[x][y] {
		return
	}
	used[x][y] = true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] == '#' {
				nx := x + i - 1
				ny := y + j - 1
				if nx < 0 || nx >= 9 || ny < 0 || ny >= 9 {
					continue
				}
				dfs(nx, ny)
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	a--
	b--
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &s[i])
	}

	dfs(a, b)

	cnt := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if used[i][j] {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
