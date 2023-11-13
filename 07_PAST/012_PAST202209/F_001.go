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

	var a [110]int
	var v [200010][]int
	var used [110]bool

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(in, &c)
		for j := 0; j < c; j++ {
			var x int
			fmt.Fscan(in, &x)
			v[x] = append(v[x], i)
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			used[j] = false
		}
		var d int
		fmt.Fscan(in, &d)
		for j := 0; j < d; j++ {
			var y int
			fmt.Fscan(in, &y)
			for _, id := range v[y] {
				used[id] = true
			}
		}
		mx, ans := -1, -2
		for j := 0; j < n; j++ {
			if !used[j] && mx < a[j] {
				mx = a[j]
				ans = j
			}
		}
		fmt.Fprintln(out, ans+1)
	}
}
