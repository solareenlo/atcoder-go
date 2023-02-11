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

	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		var d int
		fmt.Fscan(in, &d)
		m[d] = true
	}

	fmt.Println(len(m))
}
