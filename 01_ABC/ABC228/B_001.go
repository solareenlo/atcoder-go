package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	cnt := 0
	for b := x - 1; a[b] > 0; b = x - 1 {
		x = a[b]
		a[b] = 0
		cnt++
	}
	fmt.Println(cnt)
}
