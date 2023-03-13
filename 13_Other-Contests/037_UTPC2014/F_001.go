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
	for k := 0; k < Q; k++ {
		var S string
		fmt.Fscan(in, &S)
		cn := 0
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				if S[i] > S[j] {
					cn ^= 1
				}
			}
		}
		fmt.Fprintln(out, cn)
	}
}
