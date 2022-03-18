package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var a, b string
	fmt.Fscan(in, &n, &a, &b)

	j := 0
	for i := 0; i < n; i++ {
		if a[i] == b[j] {
			j++
		}
	}

	fmt.Println(2*n - j)
}
