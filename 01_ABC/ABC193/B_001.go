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

	res := 1 << 62
	for i := 0; i < n; i++ {
		var a, p, x int
		fmt.Fscan(in, &a, &p, &x)
		if a < x && p < res {
			res = p
		}
	}

	if res == 1<<62 {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}
