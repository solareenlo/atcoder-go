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

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	var m int
	fmt.Fscan(in, &m)
	px := make([]int, m+1)
	py := make([]int, m+1)
	dx := make([]int, m+1)
	dy := make([]int, m+1)

	px[0] = 1
	py[0] = 2
	dx[0] = 0
	dy[0] = 0
	for i := 0; i < m; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			px[i+1] = py[i]
			py[i+1] = -px[i]
			dx[i+1] = dy[i]
			dy[i+1] = -dx[i]
		} else if t == 2 {
			px[i+1] = -py[i]
			py[i+1] = px[i]
			dx[i+1] = -dy[i]
			dy[i+1] = dx[i]
		} else if t == 3 {
			var p int
			fmt.Fscan(in, &p)
			px[i+1] = -px[i]
			py[i+1] = py[i]
			dx[i+1] = 2*p - dx[i]
			dy[i+1] = dy[i]
		} else if t == 4 {
			var p int
			fmt.Fscan(in, &p)
			px[i+1] = px[i]
			py[i+1] = -py[i]
			dx[i+1] = dx[i]
			dy[i+1] = 2*p - dy[i]
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		b--
		sx, sy := 0, 0
		if px[a] == 1 {
			sx = x[b] + dx[a]
		} else if px[a] == 2 {
			sx = y[b] + dx[a]
		} else if px[a] == -1 {
			sx = -x[b] + dx[a]
		} else {
			sx = -y[b] + dx[a]
		}
		if py[a] == 1 {
			sy = x[b] + dy[a]
		} else if py[a] == 2 {
			sy = y[b] + dy[a]
		} else if py[a] == -1 {
			sy = -x[b] + dy[a]
		} else {
			sy = -y[b] + dy[a]
		}
		fmt.Fprintln(out, sx, sy)
	}
}
