package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	I := int(1e9)

	E := make([][]int, n)
	d := make([]int, n)
	for i := range d {
		d[i] = I
	}
	d[0] = 0
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		E[a] = append(E[a], b)
		E[b] = append(E[b], a)
	}

	Q := make([]int, 0)
	Q = append(Q, 0)
	for len(Q) > 0 {
		a := Q[0]
		Q = Q[1:]
		for _, b := range E[a] {
			if d[b] == I {
				d[b] = d[a] + 1
				Q = append(Q, b)
			}
		}
	}

	b := 120
	for _, a := range d {
		if a < I {
			if a > b {
				if a&1 != 0 {
					fmt.Println(119)
				} else {
					fmt.Println(b)
				}
			} else {
				fmt.Println(a)
			}
		} else {
			fmt.Println(b)
		}
	}
}
