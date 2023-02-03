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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make(map[int]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == '#' {
				a[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
}
