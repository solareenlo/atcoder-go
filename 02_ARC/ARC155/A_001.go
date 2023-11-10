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

	const N = 400040

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var n, k int
		var S string
		fmt.Fscan(in, &n, &k, &S)
		S = S + strings.Repeat(" ", len(S))
		s := strings.Split(S, "")
		for i := 0; i < n; i++ {
			s[n+i] = s[n-i-1]
		}
		flg := true
		for i := 0; i < 2*n; i++ {
			if s[i] != s[(i+k+n)%(2*n)] {
				flg = false
			}
		}
		if flg {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
