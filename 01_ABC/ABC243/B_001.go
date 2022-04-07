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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	s1, s2 := 0, 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if a[i] == b[j] {
				if i == j {
					s1++
				} else {
					s2++
				}
			}
		}
	}

	fmt.Println(s1)
	fmt.Println(s2)
}
