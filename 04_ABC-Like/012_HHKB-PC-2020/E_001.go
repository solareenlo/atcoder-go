package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([]string, h+1)
	for i := 1; i < h+1; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = " " + s[i] + " "
	}

	mod := int(1e9 + 7)
	pow2 := make([]int, h*w+1)
	pow2[0] = 1
	for i := 1; i <= h*w; i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	si := 0
	ans := 0
	u := [2002][2002]int{}
	l := [2002][2002]int{}
	for i := 1; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			if s[i][j] == '#' {
				continue
			}
			u[i][j] = u[i-1][j] + 1
			l[i][j] = l[i][j-1] + 1
			si++
		}
	}

	d := [2002][2002]int{}
	r := [2002][2002]int{}
	for i := h; i > 0; i-- {
		for j := w; j > 0; j-- {
			if s[i][j] == '#' {
				continue
			}
			d[i][j] = d[i+1][j] + 1
			r[i][j] = r[i][j+1] + 1
			x := u[i][j] + d[i][j] + l[i][j] + r[i][j] - 3
			ans += pow2[si-x] * (pow2[x] - 1) % mod
			ans %= mod
		}
	}

	fmt.Println((ans + mod) % mod)
}
