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

	N := 1000005
	a := make([]int, N)
	b := make([]int, N)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[x]++
		b[y]++
	}

	cnt := 0
	for i := 1; i < N; i++ {
		if cnt == 0 && a[i] != 0 {
			fmt.Fprint(out, i, " ")
		}
		cnt += a[i] - b[i]
		if cnt == 0 && b[i] != 0 {
			fmt.Fprintln(out, i)
		}
	}

}
