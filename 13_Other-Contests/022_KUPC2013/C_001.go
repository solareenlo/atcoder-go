package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(in, &m, &n)
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := 0; i < n; i++ {
		a[0][i] = 1 - a[0][i]
	}
	for i := 0; i < m; i++ {
		a[i][0] = 1 - a[i][0]
		a[i][n-1] = 1 - a[i][n-1]
	}
	sum := 0
	for i := 0; i < m; i++ {
		s, h := 0, 0
		for j := 0; j < n; j++ {
			if a[i][j] != 0 {
				s++
			} else {
				h++
			}
		}
		if h != 0 {
			sum += s + 1
		} else {
			sum += s - 1
		}
	}
	fmt.Println(sum)
}
