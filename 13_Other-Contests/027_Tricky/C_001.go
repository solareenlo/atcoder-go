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

	var T int
	fmt.Fscan(in, &T)
	for i := 0; i < T; i++ {
		var p, x int
		fmt.Fscan(in, &p, &x)
		if p == 2 {
			fmt.Fprintln(out, (1<<x)-1)
			continue
		}
		y, n := 0, 0
		for y != x {
			y++
			n++
			m := n
			for m%p == 0 {
				m /= p
				y--
			}
		}
		fmt.Fprintln(out, n)
	}
}
