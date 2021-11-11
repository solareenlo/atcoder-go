package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := [5]int{}
	a[1] = n
	a[2] = n

	black := (n - 2) * (n - 2)
	b := [5][200005]int{}
	for ; q > 0; q-- {
		var t, x int
		fmt.Fscan(in, &t, &x)
		if x < a[t] {
			for i := x; i <= a[t]; i++ {
				b[t][i] = a[3-t]
			}
			a[t] = x
		}
		black -= b[t][x] - 2
	}

	fmt.Println(black)
}
