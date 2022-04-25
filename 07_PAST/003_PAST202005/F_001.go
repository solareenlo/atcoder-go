package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := make([]string, n)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = "-"
	}
	l := 0
	r := n - 1
	for l <= r {
		c := "-"
		for _, cl := range A[l] {
			for _, cr := range A[r] {
				if cl == cr {
					c = string(cl)
					break
				}
			}
			if c != "-" {
				break
			}
		}
		if c != "-" {
			ans[l] = c
			ans[r] = c
		} else {
			fmt.Println(-1)
			return
		}
		l++
		r--
	}
	fmt.Println(strings.Join(ans, ""))
}
