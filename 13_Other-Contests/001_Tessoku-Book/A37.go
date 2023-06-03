package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, b int
	fmt.Fscan(in, &n, &m, &b)
	s := b * (n * m)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		s += a * m
	}
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		s += a * n
	}
	fmt.Println(s)
}
