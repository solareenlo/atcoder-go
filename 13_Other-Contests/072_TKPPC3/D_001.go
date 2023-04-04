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

	var h, w, q int
	fmt.Fscan(in, &h, &w, &q)
	q1 := make([]int, h+1)
	for i := 1; i <= h; i++ {
		fmt.Fscan(in, &q1[i])
		if (i % 2) != 0 {
			q1[i] += q1[i-1]
		} else {
			q1[i] = q1[i-1] - q1[i]
		}
	}
	q2 := make([]int, w+1)
	for i := 1; i <= w; i++ {
		fmt.Fscan(in, &q2[i])
		if (i % 2) != 0 {
			q2[i] += q2[i-1]
		} else {
			q2[i] = q2[i-1] - q2[i]
		}
	}
	for i := 1; i <= q; i++ {
		var a1, a2, a3, a4 int
		fmt.Fscan(in, &a1, &a2, &a3, &a4)
		fmt.Fprintln(out, (q1[a3]-q1[a1-1])*(q2[a4]-q2[a2-1]))
	}
}
