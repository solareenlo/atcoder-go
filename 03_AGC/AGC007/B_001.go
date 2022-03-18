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

	const N = 23333
	b := make([]int, N)
	for i := 1; i <= n; i++ {
		var p int
		fmt.Fscan(in, &p)
		b[p] += i
	}

	for i := 1; i <= n; i++ {
		fmt.Fprint(out, i*N, " ")
	}
	fmt.Fprintln(out)

	for i := 1; i <= n; i++ {
		fmt.Fprint(out, (n-i)*N+b[i], " ")
	}
	fmt.Fprintln(out)
}
