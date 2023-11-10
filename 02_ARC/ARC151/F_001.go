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

	var tc int
	fmt.Fscan(in, &tc)

	var l, r [3]int
	var solve func()
	solve = func() {
		fmt.Fscan(in, &l[0], &l[1], &l[2])
		fmt.Fscan(in, &r[0], &r[1], &r[2])
		an, bn, cn := 0, 0, 0
		as, bs, cs := 0, 0, 0
		for i := 0; i < 3; i++ {
			if l[i] > 0 && r[i] > 0 {
				if l[i] < r[i] {
					an++
					as += l[i]
				}
				if l[i] > r[i] {
					bn++
					bs += r[i]
				}
				if l[i] == r[i] {
					cn++
					cs += l[i]
				}
			}
		}
		if cn == 0 {
			if as >= bs {
				fmt.Fprintln(out, "Takahashi")
			} else {
				fmt.Fprintln(out, "Aoki")
			}
			return
		}
		if cn == 1 {
			if an == 2 {
				fmt.Fprintln(out, "Takahashi")
			} else if bn == 2 {
				fmt.Fprintln(out, "Aoki")
			} else if as-bs >= cs {
				fmt.Fprintln(out, "Takahashi")
			} else if bs-as >= cs {
				fmt.Fprintln(out, "Aoki")
			} else {
				if (as+bs+cs)%2 == 1 {
					fmt.Fprintln(out, "Takahashi")
				} else {
					fmt.Fprintln(out, "Aoki")
				}
			}
			return
		}
		if cn == 2 {
			if an == 1 {
				fmt.Fprintln(out, "Takahashi")
			} else if bn == 1 {
				fmt.Fprintln(out, "Aoki")
			} else {
				if cs%2 == 1 {
					fmt.Fprintln(out, "Takahashi")
				} else {
					fmt.Fprintln(out, "Aoki")
				}
			}
			return
		}
		if cs%2 == 0 {
			fmt.Fprintln(out, "Takahashi")
		} else {
			fmt.Fprintln(out, "Aoki")
		}
	}
	for tc > 0 {
		tc--
		solve()
	}
}
