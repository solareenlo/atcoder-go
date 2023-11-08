package main

import (
	"bufio"
	"fmt"
	"os"
)

var v [200005][]int
var vis [200005]bool

func dfs(p int) {
	vis[p] = true
	for _, i := range v[p] {
		if !vis[i] {
			dfs(i)
		}
	}
	if p != 1 {
		fmt.Printf("%d ", p)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		v[i] = make([]int, x)
		for j := range v[i] {
			fmt.Fscan(in, &v[i][j])
		}
	}
	dfs(1)
}
