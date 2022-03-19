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

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	s := 0
	for i := n - 1; i >= 0; i-- {
		s += (b[i] - (a[i]+s)%b[i]) % b[i]
	}
	fmt.Println(s)
}
