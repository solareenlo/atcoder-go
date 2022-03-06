package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s := [2]string{}
	fmt.Fscan(in, &s[0], &s[1])

	f := 0
	a := [2]string{}
	v := [2][10]int{}
	for q := 0; q < 2; q++ {
		if f == 0 {
			for t := 0; t < 2; t++ {
				a[t] = ""
				for i := 0; i < 10; i++ {
					v[t][i] = 0
				}
				for _, p := range s[t] {
					v[t][p-48]++
				}
			}
		}

		for k := 9 + q; k < 19; k++ {
			for i := 0; i < 10; i++ {
				if k < 10+i {
					x := min(v[0][i], v[1][k-i])
					if x != 0 {
						a[0] += strings.Repeat(string(i+'0'), x)
						a[1] += strings.Repeat(string(k-i+'0'), x)
						if k > 9 {
							f = 1
						} else {
							f = 0
						}
					}
					v[0][i] -= x
					v[1][k-i] -= x
				}
			}
		}
	}

	c := [2]string{}
	for t := 0; t < 2; t++ {
		for i := 0; i < 10; i++ {
			c[t] += strings.Repeat(string(i+'0'), v[t][i])
		}
		fmt.Fprintln(out, c[t]+a[t])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
