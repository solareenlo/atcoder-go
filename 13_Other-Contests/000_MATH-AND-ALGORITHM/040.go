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

	a := [1000009]int{}
	d := [1000009]int{}
	d[1] = 0
	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i])
		d[i+1] = d[i] + a[i]
	}

	b := [1000009]int{}
	var m int
	fmt.Fscan(in, &m)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
	}

	ans := 0
	for i := 2; i <= m; i++ {
		ans += abs(d[b[i]] - d[b[i-1]])
	}
	fmt.Fprintln(out, ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
