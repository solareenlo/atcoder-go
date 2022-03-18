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

	b := make([]int, n+1)
	b[1] = 1
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		c[i] = 1
	}

	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		c[y]++
		c[x]--
		b[y] |= b[x]
		if c[x] == 0 {
			b[x] = 0
		}
	}

	a := 0
	for i := 1; i <= n; i++ {
		a += b[i]
	}
	fmt.Println(a)
}
