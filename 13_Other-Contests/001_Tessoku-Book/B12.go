package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscan(in, &n)
	lx := 0.0
	rx := 100.0
	for rx-lx >= 0.0001 {
		mx := (lx + rx) / 2
		if mx*mx*mx+mx <= n {
			lx = mx
		} else {
			rx = mx
		}
	}
	fmt.Println(lx)
}
