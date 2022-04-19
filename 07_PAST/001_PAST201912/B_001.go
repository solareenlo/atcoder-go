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

	var now int
	fmt.Fscan(in, &now)
	for i := 0; i < n-1; i++ {
		var next int
		fmt.Fscan(in, &next)
		res := abs(now - next)
		if next == now {
			fmt.Fprintln(out, "stay")
		} else if next > now {
			fmt.Fprintln(out, "up", res)
		} else {
			fmt.Fprintln(out, "down", res)
		}
		now = next
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
