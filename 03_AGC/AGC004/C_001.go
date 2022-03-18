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

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([][]string, h)
	a := make([][]string, h)
	b := make([][]string, h)
	for i := 0; i < h; i++ {
		var S string
		fmt.Fscan(in, &S)
		s[i] = strings.Split(S, "")
		a[i] = make([]string, len(s[i]))
		b[i] = make([]string, len(s[i]))
		copy(a[i], s[i])
		copy(b[i], s[i])
		if i%2 == 0 {
			for j := 1; j < w; j++ {
				a[i][j] = "#"
			}
		} else {
			for j := 0; j < w-1; j++ {
				b[i][j] = "#"
			}
		}
		b[i][0] = "#"
		a[i][w-1] = "#"
	}

	for i := 0; i < h; i++ {
		fmt.Fprintln(out, strings.Join(a[i], ""))
	}
	fmt.Fprintln(out)
	for i := 0; i < h; i++ {
		fmt.Fprintln(out, strings.Join(b[i], ""))
	}
}
