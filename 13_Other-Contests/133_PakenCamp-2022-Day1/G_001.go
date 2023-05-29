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

	var n int
	fmt.Fscan(in, &n)
	d := make([]int, n)
	p := make([]int, n)
	d[0] = 0
	p[0] = -1
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &p[i])
		p[i]--
		d[i] = d[p[i]] + 1
	}
	doub := make([][]int, 20)
	for i := range doub {
		doub[i] = make([]int, n)
	}
	doub[0] = p
	for tr := 1; tr < 20; tr++ {
		for i := 0; i < n; i++ {
			w := doub[tr-1][i]
			if w == -1 {
				doub[tr][i] = -1
			} else {
				doub[tr][i] = doub[tr-1][w]
			}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		up := d[v] - d[u] - 1
		for tr := 19; tr >= 0; tr-- {
			if up&(1<<tr) != 0 {
				v = doub[tr][v]
			}
		}
		fmt.Fprintln(out, v+1)
	}
}
