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
		a[i]--
		b[i]--
	}

	ok := [202][202]bool{}
	for j := 1; j <= n; j++ {
		for i := 0; i+j <= n; i++ {
			ok[i][j] = true
			used := make([]bool, 2*n)
			for k := 0; k < n; k++ {
				if (i*2 <= a[k] && a[k] < (i+j)*2) || (i*2 <= b[k] && b[k] < (i+j)*2) {
					l := a[k]
					r := b[k]
					if l == -2 {
						l = b[k] - j
					}
					if r == -2 {
						r = a[k] + j
					}
					if r-l == j && i*2 <= l && l < i*2+j && !used[l] {
						used[l] = true
					} else {
						ok[i][j] = false
					}
				}
			}
			for k := 1; k < j; k++ {
				if ok[i][k] && ok[i+k][j-k] {
					ok[i][j] = true
				}
			}
		}
	}

	if ok[0][n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
