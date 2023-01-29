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

	var a [501]string
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		a[i] = " " + s
	}

	x, y := 1, 1
	var f [501][501]bool
	as, bs := 0, 0
	for x > 0 && y > 0 && x <= n && y <= m {
		if f[x][y] == true {
			fmt.Println(-1)
			return
		}
		f[x][y] = true
		as = x
		bs = y
		if a[x][y] == 'L' {
			y--
		} else if a[x][y] == 'R' {
			y++
		} else if a[x][y] == 'U' {
			x--
		} else {
			x++
		}
	}
	fmt.Println(as, bs)
}
