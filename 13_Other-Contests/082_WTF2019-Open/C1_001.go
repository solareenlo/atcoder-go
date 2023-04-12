package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100010

	var n int
	fmt.Fscan(in, &n)
	var x, y [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	x[0] = -1e17
	y[0] = 0
	for c := 61; c >= 0; c-- {
		p := 1 << c
		light := false
		for i := 1; i <= n; i++ {
			disx := x[i] - x[0]
			disy := (y[0] | p) - y[i]
			if disy < 0 || disy > disx {
				continue
			}
			if light == ((disx & disy) == disy) {
				light = false
			} else {
				light = true
			}
		}
		if light {
			y[0] |= p
		}
	}
	fmt.Println(x[0] + y[0])
}
