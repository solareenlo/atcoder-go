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
	var f [1 << 20]int
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		f[x] = x
	}

	res := 0
	for i := 1; i < (1 << 20); i++ {
		if f[i] == i {
			res++
		}
		for j := 0; j < 20; j++ {
			f[i|(1<<j)] |= f[i]
		}
	}
	fmt.Println(res)
}
