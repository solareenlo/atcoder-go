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

	const N = 10000010

	var n int
	fmt.Fscan(in, &n)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		fmt.Fprintf(out, "%d ", a[i]-a[i-1])
	}
}
