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

	const N = 200005

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([][]int, N)
	for i := 1; i <= n; i++ {
		var k int
		fmt.Fscan(in, &k)
		for j := 0; j < k; j++ {
			var x int
			fmt.Fscan(in, &x)
			a[i] = append(a[i], x)
		}
	}

	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, a[x][y-1])
	}
}
