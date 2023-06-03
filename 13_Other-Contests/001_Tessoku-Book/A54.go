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
	mp := make(map[string]int)
	for Q > 0 {
		Q--
		var cmd int
		var s string
		fmt.Fscan(in, &cmd, &s)
		if cmd == 1 {
			var x int
			fmt.Fscan(in, &x)
			mp[s] = x
		} else {
			fmt.Println(mp[s])
		}
	}
}
