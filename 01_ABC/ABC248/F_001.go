package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, P int
	fmt.Scan(&n, &P)

	f := [3010][3010][2]int{}
	f[1][1][0] = 1
	for i := 1; i <= n; i++ {
		f[i][0][1] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i+1; j++ {
			f[i][j][1] = ((f[i-1][j][1]+3*f[i-1][j-1][1]%P)%P + f[i-1][j][0]) % P
			if j-2 >= 0 {
				f[i][j][0] = (2*f[i-1][j-2][1]%P + f[i-1][j-1][0]) % P
			}
		}
	}

	for i := 1; i < n; i++ {
		fmt.Fprint(out, f[n][i][1], " ")
	}
}
