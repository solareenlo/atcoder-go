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

	const N = 200005

	var n, q int
	fmt.Fscan(in, &n, &q)
	se := make([]map[int]bool, N)
	for i := range se {
		se[i] = make(map[int]bool)
	}
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		se[i][x] = true
	}
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		if len(se[a]) != 0 {
			if len(se[a]) > len(se[b]) {
				se[a], se[b] = se[b], se[a]
			}
		}
		for k, v := range se[a] {
			if v {
				se[b][k] = true
			}
		}
		se[a] = make(map[int]bool)
		fmt.Fprintln(out, len(se[b]))
	}
}
