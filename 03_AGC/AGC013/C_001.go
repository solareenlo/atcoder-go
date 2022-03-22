package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l, t int
	fmt.Fscan(in, &n, &l, &t)

	f := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if y^1 != 0 {
			x -= t
		} else {
			x += t
		}
		p = ((p+int(math.Floor(float64(x)/float64(l))))%n + n) % n
		f[i] = (x%l + l) % l
	}
	sort.Ints(f)

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, f[(p+i)%n])
	}
}
