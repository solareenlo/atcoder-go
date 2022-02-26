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
		a[i] = 1 << 60
	}

	c := [501][501]int{}
	d := [501][501]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &c[i][j])
			a[i] = min(a[i], c[i][j])
			d[i][j] = c[i][j] - c[i][0]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if d[i][j] != d[0][j] {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
	for i := 0; i < n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
	fmt.Fprintln(out)

	for i := 0; i < n; i++ {
		fmt.Fprint(out, c[0][i]-a[0], " ")
	}
	fmt.Fprintln(out)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
