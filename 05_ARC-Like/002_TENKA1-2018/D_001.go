package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	var m int
	for m = 1; m*(m-1)/2 < n; m++ {
	}
	if m*(m-1)/2 != n {
		fmt.Fprintln(out, "No")
		return
	}

	A := make([][]int, m)
	for i := range A {
		A[i] = make([]int, m)
	}
	e := 0
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			A[i][j] = e
			A[j][i] = e
			e++
		}
	}

	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, m)
	for i := 0; i < m; i++ {
		fmt.Fprint(out, m-1)
		for j := 0; j < m; j++ {
			if j != i {
				fmt.Fprint(out, " ", A[i][j]+1)
			}
		}
		fmt.Fprintln(out)
	}
}
