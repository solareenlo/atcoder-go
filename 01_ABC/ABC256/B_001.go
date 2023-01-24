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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	x, res := 0, 0
	for i := n; i > 0; i-- {
		x += a[i]
		if x >= 4 {
			res++
		}
	}

	fmt.Println(res)
}
