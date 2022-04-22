package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	idx int
	L   = make([]int, 400040)
	R   = make([]int, 400040)
	E   = make([][]int, 150005)
)

func euler(cu, pa int) {
	L[cu] = idx
	idx++
	for _, to := range E[cu] {
		if to != pa {
			euler(to, cu)
		}
	}
	R[cu] = idx
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	root := -1
	for i := 0; i < N; i++ {
		var p int
		fmt.Fscan(in, &p)
		p--
		if p < 0 {
			root = i
		} else {
			E[p] = append(E[p], i)
			E[i] = append(E[i], p)
		}
	}
	euler(root, -1)

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if L[b] <= L[a] && R[a] <= R[b] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
