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

	var a [2002][2002]int
	var bi, bj [2002]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &a[i][j])
			bi[i] += a[i][j]
			bj[j] += a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprintf(out, "%d ", bi[i]+bj[j]-a[i][j])
		}
		fmt.Fprintln(out)
	}
}
