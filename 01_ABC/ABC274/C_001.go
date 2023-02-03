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

	par := make([]int, n*2+2)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		par[i*2] = a
		par[i*2+1] = a
	}

	ans := make([]int, n*2+2)
	for i := 2; i <= n*2+1; i++ {
		ans[i] = ans[par[i]] + 1
	}
	for i := 1; i <= n*2+1; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
