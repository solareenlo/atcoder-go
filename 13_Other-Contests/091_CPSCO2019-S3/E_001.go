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
	x, y := 0, 0
	var s [30][2]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		y ^= x
		r := 0
		for j := 0; j < 30; j++ {
			for k := 0; k < 2; k++ {
				s[j][k] += ((x >> j) & 1) ^ k
			}
			r += s[j][(y>>j)&1] * (1 << j)
		}
		fmt.Fprintln(out, r)
	}
}
