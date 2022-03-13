package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	d := make([]int, n)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}
	sort.Ints(d)

	o := make([]int, n)
	for i := 1; i < n-1; i++ {
		if d[i] < 0 {
			o[i] = d[n-1]
			d[n-1] -= d[i]
		} else {
			o[i] = d[0]
			d[0] -= d[i]
		}
	}

	fmt.Fprintln(out, d[n-1]-d[0])
	for i := 1; i < n-1; i++ {
		fmt.Fprintln(out, o[i], d[i])
	}
	fmt.Fprintln(out, d[n-1], d[0])
}
