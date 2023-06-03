package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	G := make([][]int, N)
	for i := 0; i < M; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		A--
		B--
		G[A] = append(G[A], B)
		G[B] = append(G[B], A)
	}
	d := make([]int, N)
	for i := range d {
		d[i] = -1
	}
	q := make([]int, 0)
	q = append(q, 0)
	d[0] = 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for _, i := range G[v] {
			if d[i] != -1 {
				continue
			}
			d[i] = d[v] + 1
			q = append(q, i)
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(d[i])
	}
}
