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

	type pair struct {
		x, y int
	}

	var solve func()
	solve = func() {
		var r int
		fmt.Fscan(in, &r)
		ds := make([]pair, 0)
		for d := 2; d*d <= r; d++ {
			if r%d != 0 {
				continue
			}
			e := 0
			for r%d == 0 {
				r /= d
				e++
			}
			ds = append(ds, pair{d, e})
		}
		if r != 1 {
			ds = append(ds, pair{r, 1})
		}
		ans := 1
		for _, tmp := range ds {
			p, e := tmp.x, tmp.y
			if p%4 == 1 {
				continue
			}
			if p%4 == 3 && e%2 == 1 {
				fmt.Fprintln(out, "inf")
				return
			}
			for e > 0 {
				e--
				ans *= p
			}
		}
		fmt.Fprintln(out, ans)
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		solve()
	}
}
