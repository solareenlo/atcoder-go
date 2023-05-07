package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	g := make([][]int, n+1)
	l := make([]int, n+1)
	v := make([]bool, n+1)
	d := make([]int, 0)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], i)
		g[b] = append(g[b], i)
		if a == i || b == i {
			l[i] = 1
		}
	}

	cnt := 0
	var f func(int)
	f = func(u int) {
		if v[u] {
			return
		}
		v[u] = true
		cnt++
		for _, w := range g[u] {
			f(w)
		}
		d = append(d, u)
	}

	for i := 1; i <= n; i++ {
		if l[i] != 0 {
			f(i)
		}
	}

	if cnt < n {
		fmt.Println(-1)
		return
	}

	for _, i := range d {
		fmt.Println(i)
	}
}
