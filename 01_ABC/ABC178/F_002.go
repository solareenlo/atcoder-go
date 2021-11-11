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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	b := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	j, bef := 0, -1
	for i := 0; i < n; i++ {
		if bef != a[i] {
			j = 0
		}
		if a[i] == b[i] {
			for ; j < n; j++ {
				if a[j] != a[i] && b[j] != a[i] {
					b[i], b[j] = b[j], b[i]
					break
				}
			}
			if a[i] == b[i] {
				fmt.Fprintln(out, "No")
				return
			}
		}
		bef = a[i]
	}

	fmt.Fprintln(out, "Yes")
	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, b[i])
	}
	fmt.Fprintln(out)
}
