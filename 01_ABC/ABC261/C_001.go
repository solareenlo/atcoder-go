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

	mp := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		k := mp[s]
		mp[s]++
		fmt.Fprint(out, s)
		if k > 0 {
			fmt.Fprintf(out, "(%d)", k)
		}
		fmt.Fprintln(out)
	}
}
