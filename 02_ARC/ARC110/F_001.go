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

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	fmt.Fprintln(out, (n+99)*1000)

	for j := 0; j < 1000; j++ {
		for i := 0; i < n; i++ {
			fmt.Fprintln(out, i)
		}
		for i := 0; i < 99; i++ {
			fmt.Fprintln(out, 0)
		}
	}
}
