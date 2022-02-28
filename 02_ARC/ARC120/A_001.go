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

	sum := 0
	tmp := 0
	maxi := 0
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		maxi = max(maxi, x)
		tmp += x
		sum += tmp
		fmt.Fprintln(out, sum+i*maxi)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
