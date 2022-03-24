package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	is := [2]int{}
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		is[x&1] = 1
	}

	if is[1] != 0 {
		fmt.Println(1 << (n - 1))
	} else {
		if m == 0 {
			fmt.Println((1 << n))
		} else {
			fmt.Println(0)
		}
	}
}
