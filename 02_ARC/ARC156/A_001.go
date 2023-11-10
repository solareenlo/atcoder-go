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

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var n int
		var s string
		fmt.Fscan(in, &n, &s)
		t1, t2, c0, c1 := 0, 0, 0, 0
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				c0++
			} else {
				if c1 == 0 {
					t1 = i
				}
				t2 = i
				c1++
			}
		}
		if (c1 & 1) != 0 {
			fmt.Fprintln(out, -1)
		} else if c1 > 2 || c1 == 0 || t2-t1 > 1 {
			fmt.Fprintln(out, c1>>1)
		} else if c0 > 1 {
			if t1 > 1 || t2 < n-2 {
				fmt.Fprintln(out, 2)
			} else {
				fmt.Fprintln(out, 3)
			}
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
