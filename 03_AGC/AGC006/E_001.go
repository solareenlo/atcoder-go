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

	a := [4][100001]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	p := [100001]int{}
	q := [100001]int{}
	s := [2]int{}
	for i := 1; i <= n; i++ {
		if a[2][i]%3 == 2 && abs(a[1][i]-a[2][i]) == 1 && abs(a[2][i]-a[3][i]) == 1 {
			p[i] = a[2][i]/3 + 1
			q[p[i]] = i
			if (i-p[i])&1 != 0 {
				fmt.Println("No")
				return
			}
			if a[1][i] > a[2][i] {
				s[i&1] ^= 1
			}
		} else {
			fmt.Println("No")
			return
		}
	}

	for i := 1; i <= n; i++ {
		if i != q[i] {
			q[p[i]] = q[i]
			p[i] ^= p[q[i]]
			p[q[i]] ^= p[i]
			p[i] ^= p[q[i]]
			s[i&1^1] ^= 1
		}
	}
	if s[0] != 0 || s[1] != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
