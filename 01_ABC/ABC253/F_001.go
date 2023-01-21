package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	op, x, y, z int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	const N = 200020
	var a [N]node
	var las [N]int
	g := make([][]int, N)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &a[i].op, &a[i].x, &a[i].y)
		if a[i].op == 1 {
			fmt.Fscan(in, &a[i].z)
		} else if a[i].op == 2 {
			las[a[i].x] = i
		} else {
			if las[a[i].x] == 0 {
				a[i].x = 0
			} else {
				g[las[a[i].x]] = append(g[las[a[i].x]], i)
			}
		}
	}

	tr := make([]int, N)
	for i := 1; i <= q; i++ {
		for _, j := range g[i] {
			sum := 0
			for k := a[j].y; k > 0; k -= lowbit(k) {
				sum += tr[k]
			}
			a[j].x = a[i].y - ask(a[j].y, tr)
		}
		if a[i].op == 1 {
			add(a[i].x, a[i].z, m, tr)
			add(a[i].y+1, -a[i].z, m, tr)
		}
		if a[i].op == 3 {
			fmt.Fprintln(out, a[i].x+ask(a[i].y, tr))
		}
	}
}

func lowbit(x int) int { return x & (-x) }

func add(x, y, m int, tr []int) {
	for i := x; i <= m; i += lowbit(i) {
		tr[i] += y
	}
}

func ask(x int, tr []int) int {
	sum := 0
	for i := x; i > 0; i -= lowbit(i) {
		sum += tr[i]
	}
	return sum
}
