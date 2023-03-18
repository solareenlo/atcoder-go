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
	a := make([]int, n+1)
	b := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, a[i]+b[i])
	}
}
