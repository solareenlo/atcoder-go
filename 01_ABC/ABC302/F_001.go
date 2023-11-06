package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const N = 400400

	var d [N]int
	var e [N][]int

	var m, n int
	fmt.Fscan(in, &m, &n)
	for i := 1; i <= m; i++ {
		var t int
		fmt.Fscan(in, &t)
		for t > 0 {
			t--
			var x int
			fmt.Fscan(in, &x)
			e[i+n] = append(e[i+n], x)
			e[x] = append(e[x], i+n)
		}
	}
	for i := range d {
		d[i] = INF
	}
	q := make([]int, 0)
	q = append(q, 1)
	d[1] = 0
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range e[x] {
			if d[y] > d[x]+1 {
				d[y] = d[x] + 1
				q = append(q, y)
			}
		}
	}
	if d[n] > N {
		fmt.Println(-1)
	} else {
		fmt.Println(d[n]/2 - 1)
	}
}
