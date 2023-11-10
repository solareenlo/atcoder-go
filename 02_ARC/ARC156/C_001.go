package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 5005

	var e [N][]int
	var d, p [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
		d[u]++
		d[v]++
	}
	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			q = append(q, i)
		}
	}
	l := 0
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if l != 0 {
			p[x] = l
			p[l] = x
			l = 0
		} else {
			l = x
		}
		for _, y := range e[x] {
			d[y]--
			if d[y] == 1 {
				q = append(q, y)
			}
		}
	}
	if l != 0 {
		p[l] = l
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", p[i])
	}
}
