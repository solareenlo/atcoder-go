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

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := make([]int, q+1)
	r := make([]int, q+1)
	x := make([]int, q+1)
	query := make([][]int, 400004)
	for i := 1; i < q+1; i++ {
		fmt.Fscan(in, &l[i], &r[i], &x[i])
		query[r[i]] = append(query[r[i]], i)
	}

	p := [60]int{}
	d := [60]int{}
	for i := 1; i < n+1; i++ {
		pp := i
		for j := 59; j >= 0; j-- {
			if a[i]>>j&1 != 0 {
				if p[j] == 0 {
					d[j] = a[i]
					p[j] = pp
					continue
				}
				if pp > p[j] {
					pp, p[j] = p[j], pp
					d[j], a[i] = a[i], d[j]
				}
				a[i] ^= d[j]
			}
		}
		for _, id := range query[i] {
			for j := 59; j >= 0; j-- {
				if (x[id]>>j&1 != 0) && (l[id] <= p[j]) {
					x[id] ^= d[j]
				}
			}
		}
	}

	for i := 1; i < q+1; i++ {
		if x[i] != 0 {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
}
