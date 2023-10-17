package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	queue []int
}

func (q *Queue) Push(x int) {
	q.queue = append(q.queue, x)
}

func (q *Queue) Pop() int {
	x := q.queue[0]
	q.queue = q.queue[1:]
	return x
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	C := make([]int, n)
	Z := make([]int, n)
	G := make([][]int, n)

	for i := 0; i < m; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		p--
		q--
		C[p]++
		G[q] = append(G[q], p)
	}

	queue := Queue{}

	for u := 0; u < n; u++ {
		if C[u] == 0 {
			Z[u] = -1
			queue.Push(u)
		}
	}

	for !queue.IsEmpty() {
		u := queue.Pop()
		for _, v := range G[u] {
			if Z[v] == 0 {
				C[v]--
				if C[v] != 0 && Z[u] != -1 {
					continue
				}
				queue.Push(v)
				if Z[u] == -1 {
					Z[v] = 1
				} else {
					Z[v] = -1
				}
			}
		}
	}

	if Z[0] != 0 {
		if Z[0] == 1 {
			fmt.Println("Snuke")
		} else {
			fmt.Println("Sothe")
		}
	} else {
		fmt.Println("Draw")
	}
}
