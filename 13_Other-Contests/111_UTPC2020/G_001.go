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
	var a [10][10]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	fmt.Fprintln(out, 1<<n)
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				t := ((i >> j) & 1) ^ ((i >> k) & 1)
				if t != 0 {
					fmt.Fprintf(out, "%d ", -a[j][k])
				} else {
					fmt.Fprintf(out, "%d ", a[j][k])
				}
			}
			fmt.Fprintln(out)
		}
	}
}
