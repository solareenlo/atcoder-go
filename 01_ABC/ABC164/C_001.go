package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	m := map[string]struct{}{}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		m[s] = struct{}{}
	}

	fmt.Println(len(m))
}
