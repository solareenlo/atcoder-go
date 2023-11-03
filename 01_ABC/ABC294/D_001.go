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

	var come [500005]bool

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	recall := 1
	call := 0
	for Q > 0 {
		Q--
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			call++
		}
		if op == 2 {
			var tmp int
			fmt.Fscan(in, &tmp)
			come[tmp] = true
		}
		if op == 3 {
			for come[recall] {
				recall++
			}
			fmt.Fprintln(out, recall)
		}
	}
}
