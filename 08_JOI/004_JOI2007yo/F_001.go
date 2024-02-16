package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, n int
	fmt.Fscan(in, &a, &b, &n)
	var v [20][20]int
	for n > 0 {
		n--
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[y][x] = 1
	}
	var d [20][20]int
	d[0][1] = 1
	for i := 1; i <= b; i++ {
		for j := 1; j <= a; j++ {
			if v[i][j] != 0 {
				continue
			}
			d[i][j] = d[i-1][j] + d[i][j-1]
		}
	}
	fmt.Println(d[b][a])
}
