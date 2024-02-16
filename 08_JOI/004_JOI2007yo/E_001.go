package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B, C, n int
	fmt.Fscan(in, &A, &B, &C, &n)
	s := make([]int, A+B+C+1)
	for i := range s {
		s[i] = 2
	}
	var c [1000][4]int
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fscan(in, &c[i][j])
		}
		if c[i][3] != 0 {
			for j := 0; j < 3; j++ {
				s[c[i][j]] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		if c[i][3] == 0 {
			for j := 0; j < 3; j++ {
				f := true
				for k := 0; k < 3; k++ {
					if f && (k == j || s[c[i][k]] == 1) {
						f = true
					} else {
						f = false
					}
				}
				if f {
					s[c[i][j]] = 0
				}
			}
		}
	}
	for i := 0; i < A+B+C; {
		i++
		fmt.Println(s[i])
	}
}
