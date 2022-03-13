package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 10005

var (
	e  = make([][]int, N)
	cc int
	a  = make([]int, N)
	c  = make([]int, N)
	s  int
)

func dfs(x, y int) {
	for _, i := range e[x] {
		if i != y {
			dfs(i, x)
		}
	}
	a[x] = c[cc]
	cc++
	s += a[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	tmp := c[:n]
	sort.Ints(tmp)

	dfs(1, 0)

	fmt.Fprintln(out, s-a[1])
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
}
