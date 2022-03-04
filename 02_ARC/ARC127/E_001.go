package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)

	top, cnt := 0, 0
	st := make([]int, 5050)
	for i := 1; i <= a+b; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			top++
			cnt++
			st[top] = cnt
		} else {
			top--
		}
	}

	f := [5050][5050]int{}
	for i := 0; i <= a; i++ {
		f[0][i] = 1
	}

	const mod = 998244353
	for i := 1; i <= top; i++ {
		for j := 1; j <= a; j++ {
			if j <= st[i] {
				f[i][j] = (f[i][j-1] + f[i-1][j-1]) % mod
			} else {
				f[i][j] = f[i][j-1]
			}
		}
	}
	fmt.Println(f[top][a])
}
