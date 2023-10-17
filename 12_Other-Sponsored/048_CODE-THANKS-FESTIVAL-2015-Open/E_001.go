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

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s string
		var a [26]bool
		var c, d string
		fmt.Fscan(in, &c, &d)
		for j := 0; j < len(d); j++ {
			a[d[j]-'a'] = true
		}
		for j := 0; j < len(c); j++ {
			if a[c[j]-'a'] {
				s += string(c[j])
			}
		}
		if strings.Index(s, d) == -1 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
