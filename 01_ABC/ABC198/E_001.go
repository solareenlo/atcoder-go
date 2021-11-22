package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	c   = [100010]int{}
	s   = [100010]int{}
	res = [100010]int{}
	e   = make([][]int, 100010)
)

func dfs(u, p int) {
	if s[c[u]] == 0 {
		res[u] = 1
	}
	s[c[u]]++
	for _, x := range e[u] {
		if x != p {
			dfs(x, u)
		}
	}
	s[c[u]]--
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &c[i])
	}

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}

	dfs(1, 0)

	for i := 1; i < n+1; i++ {
		if res[i] != 0 {
			fmt.Fprintln(out, i)
		}
	}
}
