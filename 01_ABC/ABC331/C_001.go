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

	var a [1000003]int
	var b [200005]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		a[b[i]]++
	}
	for i := 1000001; i > 0; i-- {
		a[i] *= i
		a[i] += a[i+1]
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", a[b[i]+1])
	}
}
