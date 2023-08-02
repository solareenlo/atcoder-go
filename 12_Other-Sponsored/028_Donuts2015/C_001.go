package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, st [100010]int

	var n int
	fmt.Fscan(in, &n)
	top := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
		fmt.Println(top)
		for top > 0 && st[top] < h[i] {
			top--
		}
		top++
		st[top] = h[i]
	}
}
