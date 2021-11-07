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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum ^= a[i]
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = sum ^ a[i]
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, b[i])
		if i != n-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
