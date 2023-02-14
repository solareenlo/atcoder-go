package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 20020
	const INF = int(1e18)

	var m, n int
	fmt.Fscan(in, &m, &n)

	w := make([]int, N)
	for i := range w {
		w[i] = INF
	}
	p := make([]int, N)
	for i := range p {
		p[i] = -1
	}
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		w[y-x+300] = min(w[y-x+300], x)
	}

	p[1] = 1
	q := make([]int, 0)
	q = append(q, 1)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for i := -300; i < 300; i++ {
			if x < w[i+300] || x+i < 0 || x+i >= 20000 || p[x+i] != -1 {
				continue
			}
			p[x+i] = p[x] + 1
			q = append(q, x+i)
		}
	}
	fmt.Println(p[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
