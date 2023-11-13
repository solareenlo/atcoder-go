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

	var n, m int
	fmt.Fscan(in, &n, &m)
	vertex := make(map[pair]bool)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if u > v {
			u, v = v, u
		}
		if u == v || n < v {
			continue
		}
		vertex[pair{u, v}] = true
	}
	if len(vertex) == m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
