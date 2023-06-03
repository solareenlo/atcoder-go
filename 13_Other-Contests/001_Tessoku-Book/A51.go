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

	var Q int
	fmt.Fscan(in, &Q)
	st := make([]string, 0)
	for Q > 0 {
		Q--
		var cmd int
		fmt.Fscan(in, &cmd)
		if cmd == 1 {
			var s string
			fmt.Fscan(in, &s)
			st = append(st, s)
		} else if cmd == 2 {
			fmt.Fprintln(out, st[len(st)-1])
		} else {
			st = st[:len(st)-1]
		}
	}
}
