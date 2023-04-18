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

	const MX = 100005

	var n int
	fmt.Fscan(in, &n)

	var a [MX]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var d [MX]int
	for i := 0; i < n; i++ {
		var b int
		fmt.Fscan(in, &b)
		d[b]++
	}

	ans := MX
	var c [MX]int
	for i := 0; i < n; i++ {
		c[a[i]]++
		ans = min(ans, d[a[i]]/c[a[i]])
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
