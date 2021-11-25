package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	N := 808
	s := make([][]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			s[i] = make([]int, N)
		}
		s[i][0] = 0
		s[0][i] = 0
	}

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	lim := ((k * k) / 2) + 1
	ng, ok := -1, 1000000000
	for ng+1 < ok {
		mid := (ng + ok) / 2
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j]
				if a[i][j] > mid {
					s[i+1][j+1]++
				}
			}
		}
		ext := false
		for i := 0; i < n-k+1; i++ {
			for j := 0; j < n-k+1; j++ {
				if (s[i+k][j+k] + s[i][j] - s[i][j+k] - s[i+k][j]) < lim {
					ext = true
				}
			}
		}
		if ext {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
}
