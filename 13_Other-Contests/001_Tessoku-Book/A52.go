package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(in, &Q)
	q := make([]string, 0)
	for Q > 0 {
		Q--
		var cmd int
		fmt.Fscan(in, &cmd)
		if cmd == 1 {
			var s string
			fmt.Fscan(in, &s)
			q = append(q, s)
		} else if cmd == 2 {
			fmt.Println(q[0])
		} else {
			q = q[1:]
		}
	}
}
