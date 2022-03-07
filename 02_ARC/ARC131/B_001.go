package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	A := make([]string, h+2)
	A[0] = strings.Repeat(" ", w+2)
	for i := 1; i <= h; i++ {
		fmt.Fscan(in, &A[i])
		A[i] = " " + A[i] + " "
	}
	A[h+1] = strings.Repeat(" ", w+2)

	a := make([][]string, h+2)
	for i := range A {
		a[i] = strings.Split(A[i], "")
	}

	dx := [5]int{1, -1, 0, 0}
	dy := [5]int{0, 0, 1, -1}
	vis := make([]bool, 1000)
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if a[i][j] == "." {
				for k := 0; k < 4; k++ {
					vis[a[i+dx[k]][j+dy[k]][0]] = true
				}
				for k := 0; k < 5; k++ {
					if !vis[k+'1'] {
						a[i][j] = string(k + '1')
					} else {
						vis[k+'1'] = false
					}
				}
			}
		}
	}

	for i := 1; i <= h; i++ {
		tmp := strings.Join(a[i][1:w+1], "")
		fmt.Fprintln(out, tmp)
	}
}
