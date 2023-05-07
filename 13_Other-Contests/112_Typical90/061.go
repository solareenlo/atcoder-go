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

	var Q int
	fmt.Fscan(in, &Q)

	t := make([]int, Q)
	x := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &t[i], &x[i])
	}

	u := make([]int, 0)
	v := make([]int, 0)
	m := 0
	for i := 0; i < Q; i++ {
		if t[i] == 1 {
			u = append(u, x[i])
			m++
		} else if t[i] == 2 {
			v = append(v, x[i])
		} else {
			if x[i] > m {
				fmt.Fprintln(out, v[x[i]-m-1])
			} else {
				fmt.Fprintln(out, u[m-x[i]])
			}
		}
	}
}
