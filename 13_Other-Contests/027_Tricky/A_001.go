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
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a == -9223372036854775808 && b == -1 {
			fmt.Fprintln(out, "9223372036854775808")
		} else {
			fmt.Fprintln(out, a/b)
		}
	}
}
