package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 200200

	c := make([]map[int]bool, 400400)
	for i := range c {
		c[i] = make(map[int]bool)
	}

	var n, m, h, k int
	fmt.Fscan(in, &n, &m, &h, &k)
	var s string
	fmt.Fscan(in, &s)
	s = " " + s + " "
	for i := 1; i <= m; i++ {
		var xx, yy int
		fmt.Fscan(in, &xx, &yy)
		c[xx+M][yy+M] = true
	}
	x := M
	y := M
	for i := 1; i <= n; i++ {
		if s[i] == 'R' {
			x++
		}
		if s[i] == 'L' {
			x--
		}
		if s[i] == 'U' {
			y++
		}
		if s[i] == 'D' {
			y--
		}
		h--
		if h < 0 {
			fmt.Println("No")
			return
		}
		if c[x][y] && h <= k {
			h = k
			c[x][y] = false
		}
	}
	fmt.Println("Yes")
}
