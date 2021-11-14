package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	bit  = [300003]int{}
	n, q int
)

func add(i, x int) {
	for i <= n {
		bit[i] ^= x
		i += i & -i
	}
}

func sum(i int) int {
	res := 0
	for i > 0 {
		res ^= bit[i]
		i -= i & -i
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &q)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		add(i, a)
	}
	for ; q > 0; q-- {
		var t, x, y int
		fmt.Fscan(in, &t, &x, &y)
		if t == 1 {
			add(x, y)
		} else {
			fmt.Println(sum(y) ^ sum(x-1))
		}
	}
}
