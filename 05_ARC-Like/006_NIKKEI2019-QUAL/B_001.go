package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var a, b, c string
	fmt.Fscan(in, &n, &a, &b, &c)

	m := 0
	for i := 0; i < n; i++ {
		if a[i] == b[i] && b[i] == c[i] {
			m += 0
		} else if a[i] == b[i] || b[i] == c[i] || a[i] == c[i] {
			m += 1
		} else {
			m += 2
		}
	}
	fmt.Println(m)
}
