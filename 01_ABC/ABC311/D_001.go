package main

import (
	"bufio"
	"fmt"
	"os"
)

var d, a [200][200]bool
var S [200]string

func dfs(i, j int) {
	a[i][j] = true
	di := []int{-1, 0, 1, 0}
	dj := []int{0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		ni := i
		nj := j
		for {
			d[ni][nj] = true
			if S[ni+di[k]][nj+dj[k]] == '#' {
				break
			}
			ni += di[k]
			nj += dj[k]
		}
		if !a[ni][nj] {
			dfs(ni, nj)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}
	dfs(1, 1)
	ans := 0
	for i := 0; i < 200-1; i++ {
		for j := 0; j < 200-1; j++ {
			if d[i][j] {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
