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

	var q int
	fmt.Fscan(in, &q)

	type pair struct{ x, y int }
	st := make([]pair, 200005)
	e := 0
	b := 0
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var c string
			var x int
			fmt.Fscan(in, &c, &x)
			st[e] = pair{int(c[0] - 'a'), x}
			e++
		} else {
			var x int
			fmt.Fscan(in, &x)
			del := make([]int, 26)
			for b < e {
				if st[b].y >= x {
					del[st[b].x] += x
					st[b].y -= x
					break
				} else {
					x -= st[b].y
					del[st[b].x] += st[b].y
					b++
				}
			}
			ans := 0
			for _, g := range del {
				ans += g * g
			}
			fmt.Fprintln(out, ans)
		}
	}
}
