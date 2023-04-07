package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var x, w [100100]float64
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &w[i])
	}
	l, r := -1e7, 1e7
	for i := 0; i < 100; i++ {
		m := (l + r) / 2.0
		d := -1.0
		var k int
		for j := 0; j < n; j++ {
			e := math.Abs(m-x[j]) * w[j]
			if d < e {
				d = e
				k = j
			}
		}
		if x[k] < m {
			r = m
		} else {
			l = m
		}
	}
	fmt.Println(l)
}
