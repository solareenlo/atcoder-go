package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var c [2020][]int
	var v [2020][2020]bool
	var v1 [2020]bool

	var n, m int
	fmt.Fscan(in, &n, &m)
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		c[x] = append(c[x], y)
		v[x][y] = true
	}
	ans := 0
	for x := 1; x <= n; x++ {
		q := make([]int, 0)
		for i := 1; i <= n; i++ {
			v1[i] = false
		}
		q = append(q, x)
		for len(q) > 0 {
			y := q[0]
			q = q[1:]
			if v1[y] {
				continue
			}
			v1[y] = true
			if !v[x][y] && x != y {
				ans++
			}
			for i := 0; i < len(c[y]); i++ {
				q = append(q, c[y][i])
			}
		}
	}
	fmt.Println(ans)
}
