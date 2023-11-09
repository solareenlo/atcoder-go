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

	const N = 200005

	var p, v, a [N]int
	var vis [N]bool

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		v[p[i]]++
	}
	for q > 0 {
		q--
		var m int
		fmt.Fscan(in, &m)
		ans := 0
		for i := 1; i <= m; i++ {
			fmt.Fscan(in, &a[i])
			ans += v[a[i]] + 1
			vis[a[i]] = true
		}
		for i := 1; i <= m; i++ {
			if vis[p[a[i]]] {
				ans -= 2
			}
		}
		for i := 1; i <= m; i++ {
			vis[a[i]] = false
		}
		fmt.Fprintln(out, ans)
	}
}
