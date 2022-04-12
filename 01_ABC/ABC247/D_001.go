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

	var n int
	fmt.Fscan(in, &n)

	t := 0
	s := 1
	const N = 1000007
	val := make([]int, N)
	cnt := make([]int, N)
	for i := 1; i <= n; i++ {
		var op int
		fmt.Fscan(in, &op)
		var x, y int
		if op == 1 {
			fmt.Fscan(in, &x, &y)
			t++
			val[t] = x
			cnt[t] = y
		} else {
			fmt.Fscan(in, &x)
			y = 0
			for x > cnt[s] {
				y += val[s] * cnt[s]
				x -= cnt[s]
				s++
			}
			cnt[s] -= x
			y += val[s] * x
			fmt.Fprintln(out, y)
		}
	}
}
