package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [200200]int
	var M, E, X [1000]int

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var s string
	fmt.Fscan(in, &s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' {
			M[1<<a[i]]++
		} else if s[i] == 'E' {
			for j := 1; j < 8; j++ {
				E[1<<a[i]|j] += M[j]
			}
		} else {
			for j := 1; j < 8; j++ {
				X[1<<a[i]|j] += E[j]
			}
		}
	}
	fmt.Println(X[7]*3 + X[3]*2 + X[1] + X[5])
}
