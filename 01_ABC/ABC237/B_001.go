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

	a := make([]int, h*w+1)
	for i := 1; i <= h*w; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 1; i <= w; i++ {
		for j := i; j <= h*w; j += w {
			fmt.Fprint(out, a[j], " ")
		}
		fmt.Fprintln(out)
	}
}
