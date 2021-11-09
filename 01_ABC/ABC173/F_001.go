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

	res := n * (n + 1) * (n + 2) / 6
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if u > v {
			u, v = v, u
		}
		res -= u * (n - v + 1)
	}
	fmt.Println(res)
}
