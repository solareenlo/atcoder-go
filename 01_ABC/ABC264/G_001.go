package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const x = 729
	const t = 27
	var c [19683]int
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		k := 0
		for j := range s {
			k = t*k + int(s[j]) - 96
		}
		fmt.Fscan(in, &c[k])
	}
	e := make([][]pair, x)
	for i := 1; i < t; i++ {
		e[0] = append(e[0], pair{i, -c[i]})
	}
	for i := 1; i < x; i++ {
		for k := 1; k < t; k++ {
			n = i%t*t + k
			tmp := 0
			if i >= t {
				tmp = c[t*i+k]
			}
			e[i] = append(e[i], pair{n, -c[k] - c[n] - tmp})
		}
	}
	ans := 4_557_430_888_798_830_399
	d := make([]int, x)
	for i := range d {
		d[i] = ans
	}
	d[0] = 0
	for i := 0; i < x; i++ {
		for k := 0; k < x; k++ {
			for j := range e[k] {
				t := e[k][j].x
				z := e[k][j].y
				if d[k]+z < d[t] {
					d[t] = d[k] + z
					if i > x-2 {
						fmt.Println("Infinity")
						os.Exit(0)
					}
				}
			}
		}
	}

	for i := 1; i < x; i++ {
		ans = min(ans, d[i])
	}
	fmt.Println(-ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
