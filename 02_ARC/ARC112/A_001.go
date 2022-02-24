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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		maxi := max(r-l+1-l, 0)
		fmt.Fprintln(out, (1+maxi)*maxi/2)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
