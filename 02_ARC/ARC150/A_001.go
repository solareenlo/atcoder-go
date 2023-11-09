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

	const N = 1000000

	var a, b [N]int

	var t int
	fmt.Fscan(in, &t)
	for t1 := 1; t1 <= t; t1++ {
		var n, k int
		var s string
		fmt.Fscan(in, &n, &k, &s)
		s = "0" + s + "0"
		for i := 1; i <= n; i++ {
			if s[i] == '0' {
				a[i] = a[i-1] + 1
			} else {
				a[i] = a[i-1]
			}
		}
		for i := 1; i <= n; i++ {
			if s[i] == '1' {
				b[i] = b[i-1] + 1
			} else {
				b[i] = b[i-1]
			}
		}
		ans := 0
		for i := k; i <= n; i++ {
			if a[i] == a[i-k] && b[n] == b[i] && b[i-k] == 0 {
				ans++
			}
		}
		if ans == 1 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
