package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N, M int
	fmt.Fscan(in, &N, &M)
	w := make([][]pair, N)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		w[a] = append(w[a], pair{b, c})
		w[b] = append(w[b], pair{a, c})
	}
	p := make([]int, N)
	for i := range p {
		p[i] = int(1e16)
	}
	p[0] = 0
	q := make([]int, 0)
	q = append(q, 0)
	for len(q) > 0 {
		F := q[0]
		q = q[1:]
		for _, d := range w[F] {
			to := d.x
			dis := d.y
			if p[to] <= dis+p[F] {
				continue
			}
			p[to] = dis + p[F]
			q = append(q, to)
		}
	}
	for i := 0; i < N; i++ {
		if p[i] == int(1e16) {
			p[i] = -1
		}
		fmt.Println(p[i])
	}
}
