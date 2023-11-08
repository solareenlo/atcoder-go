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

	const N = 200200

	var a, cl [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var i, j int
	for i, j = 1, 1; cl[i] == 0; {
		cl[i] = j
		i = a[i]
		j++
	}
	fmt.Fprintln(out, j-cl[i])
	fmt.Fprintf(out, "%d ", i)
	zs := i
	i = a[i]
	for ; zs != i; i = a[i] {
		fmt.Fprintf(out, "%d ", i)
	}
}
