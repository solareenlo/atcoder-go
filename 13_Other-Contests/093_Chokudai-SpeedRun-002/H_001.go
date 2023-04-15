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

	const N = 200010

	var n int
	fmt.Fscan(in, &n)
	var a, b [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		if a[i] == b[i] {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, abs(a[i]-b[i]))
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
