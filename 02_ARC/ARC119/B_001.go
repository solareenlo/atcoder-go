package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s, t string
	fmt.Fscan(in, &n, &s, &t)

	a := make([]int, 0)
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			a = append(a, i)
		}
		if t[i] == '0' {
			b = append(b, i)
		}
	}

	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	if len(a) != len(b) {
		res = -1
	}
	fmt.Println(res)
}
