package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100010
	const INF = math.MaxInt

	var e [N][]int
	var a, b, d [N]int
	var vis [N]bool

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := range d {
		d[i] = INF
	}
	if a[1] == b[1] {
		d[1] = 1
	} else {
		d[1] = 0
	}
	q := list.New()
	q.PushBack(1)
	for q.Len() > 0 {
		x := q.Front().Value.(int)
		q.Remove(q.Front())
		if vis[x] {
			continue
		}
		vis[x] = true
		for _, y := range e[x] {
			var w int
			if a[y] == b[d[x]+1] {
				w = 1
			} else {
				w = 0
			}
			if d[y] > d[x]+w {
				d[y] = d[x] + w
				if w != 0 {
					q.PushBack(y)
				} else {
					q.PushFront(y)
				}
			}
		}
	}
	if d[n] < k {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
