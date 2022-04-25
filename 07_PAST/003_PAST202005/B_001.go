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

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)

	cnt := make([]int, 55)
	a := make([][]int, 100005)
	for ; Q > 0; Q-- {
		var x, n int
		fmt.Fscan(in, &x, &n)
		if x == 1 {
			ans := 0
			for _, i := range a[n] {
				ans += N - cnt[i]
			}
			fmt.Fprintln(out, ans)
		} else {
			var m int
			fmt.Fscan(in, &m)
			a[n] = append(a[n], m)
			cnt[m]++
		}
	}
}
