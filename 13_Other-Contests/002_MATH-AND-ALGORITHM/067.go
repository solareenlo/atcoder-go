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

	var h, w int
	fmt.Fscan(in, &h, &w)

	p := [2002][2002]int{}
	ta := make([]int, h)
	yo := make([]int, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var a int
			fmt.Fscan(in, &a)
			p[i][j] = a
			ta[i] += a
			yo[j] += a
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(out, ta[i]+yo[j]-p[i][j], " ")
		}
		fmt.Fprintln(out)
	}
}
