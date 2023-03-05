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

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	s := make([]string, 10001)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
	}
	c := make([]string, 10001)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &c[i])
	}

	for i := 1; i <= q; i++ {
		var k int
		for j := 1; j <= n; j++ {
			k = strings.Index(s[j], c[i])
			if k >= 0 {
				fmt.Fprintln(out, j, k+1)
				break
			} else {
				k = -1
			}
		}
		if k < 0 {
			fmt.Fprintln(out, "NA")
		}
	}
}
