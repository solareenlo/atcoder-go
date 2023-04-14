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
	x := 1
	var a, b, c [100010]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
		if b[i] == 0 {
			a[i] = x
			x++
		} else {
			a[i] = a[b[i]-1]
		}
	}
	c[0] = -1
	j := -1
	for i := 0; i < n; i++ {
		for j >= 0 && a[i] != a[j] {
			j = c[j]
		}
		j++
		c[i+1] = j
	}
	for i := 0; i < n; i++ {
		if c[i+1] != b[i] {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
	fmt.Fprintln(out)
}
