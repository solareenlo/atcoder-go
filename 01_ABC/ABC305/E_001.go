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

	const N = 200200

	type V struct {
		e []int
		d int
	}

	v := make([]V, N)
	for i := range v {
		v[i].d = -1
	}

	var q [N][]int

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x].e = append(v[x].e, y)
		v[y].e = append(v[y].e, x)
	}
	for i := 1; i <= k; i++ {
		var p, d int
		fmt.Fscan(in, &p, &d)
		q[d] = append(q[d], p)
	}
	for i := n; i >= 0; i-- {
		for _, j := range q[i] {
			if v[j].d == -1 {
				v[j].d = i
				if v[j].d != 0 {
					for _, k := range v[j].e {
						q[i-1] = append(q[i-1], k)
					}
				}
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if v[i].d >= 0 {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
	for i := 1; i <= n; i++ {
		if v[i].d >= 0 {
			fmt.Fprintf(out, "%d ", i)
		}
	}
}
