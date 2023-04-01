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

	var Q int
	fmt.Fscan(in, &Q)
	var d [200009]int
	var p [18][200009]int
	for i := 0; i < Q; i++ {
		var s string
		var a, b int
		fmt.Fscan(in, &s, &a, &b)
		a--
		b--
		if s[0] == 'A' {
			p[0][b] = a
			d[b] = d[a] + 1
			for j := 1; j <= 17; j++ {
				p[j][b] = p[j-1][p[j-1][b]]
			}
		} else {
			x := d[a] + d[b]
			if d[a] < d[b] {
				a, b = b, a
			}
			for j := 17; j >= 0; j-- {
				if d[a]-d[b] >= 1<<j {
					a = p[j][a]
				}
			}
			for j := 17; j >= 0; j-- {
				if p[j][a] != p[j][b] {
					a = p[j][a]
					b = p[j][b]
				}
			}
			tmp := -1
			if a != b {
				tmp = 1
			}
			fmt.Fprintln(out, x-d[a]*2+tmp)
		}
	}
}
