package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	mp := make(map[pair]bool)
	for n > 0 {
		n--
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a > b {
			a, b = b, a
		}
		mp[pair{a, b}] = true
	}
	fmt.Println(len(mp))
}
