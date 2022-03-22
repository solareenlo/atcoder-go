package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	vis  = make([]bool, 200002)
	ans  = make([]int, 0)
	e    = make([][]int, 200002)
	flag bool
)

func dfs(x int) {
	vis[x] = true
	ans = append(ans, x)
	for _, y := range e[x] {
		if !vis[y] {
			dfs(y)
			if flag {
				return
			}
		}
	}
	flag = true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}

	dfs(1)
	ans = reverseOrderInt(ans)
	ans = ans[:len(ans)-1]
	flag = false
	dfs(1)

	fmt.Fprintln(out, len(ans))
	for _, i := range ans {
		fmt.Fprint(out, i, " ")
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
