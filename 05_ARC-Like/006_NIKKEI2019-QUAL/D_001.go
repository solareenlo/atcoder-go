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

	var n, m int
	fmt.Fscan(in, &n, &m)

	e := make([][]int, 100005)
	ind := make([]int, 100005)
	for i := 1; i <= n-1+m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		e[a] = append(e[a], b)
		ind[b]++
	}

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if ind[i] == 0 {
			q = append(q, i)
		}
	}

	dad := make([]int, 100005)
	for len(q) != 0 {
		a := q[0]
		q = q[1:]
		for _, i := range e[a] {
			ind[i]--
			if ind[i] == 0 {
				dad[i] = a
				q = append(q, i)
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, dad[i])
	}
}
