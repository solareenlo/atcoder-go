package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var v [55][55]bool

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := n * (n - 1) / 2
	for i := 1; i <= m; i++ {
		var _x int
		fmt.Fscan(in, &_x)
		for j := 2; j <= n; j++ {
			var x int
			fmt.Fscan(in, &x)
			if !v[_x][x] {
				ans--
			}
			v[x][_x] = true
			v[_x][x] = true
			_x = x
		}
	}
	fmt.Println(ans)
}
