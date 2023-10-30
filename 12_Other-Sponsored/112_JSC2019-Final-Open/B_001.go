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
	var a, b [300][300]bool
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < n; j++ {
			a[i][j] = (s[j] == '1')
		}
	}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < n; j++ {
			b[i][j] = (s[j] == '1')
		}
	}

	var ans [300][300]bool
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = true
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !b[i][j] {
				for k := 0; k < n; k++ {
					if a[i][k] {
						ans[k][j] = false
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] {
				ok := false
				for k := 0; k < n; k++ {
					if a[i][k] && ans[k][j] {
						ok = true
					}
				}
				if !ok {
					fmt.Fprintln(out, -1)
					return
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] {
				fmt.Fprint(out, 1)
			} else {
				fmt.Fprint(out, 0)
			}
		}
		fmt.Fprintln(out)
	}
}
