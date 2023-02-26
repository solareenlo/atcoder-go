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
	if n < 4 {
		fmt.Println(-1)
		return
	}
	var f [1000][1000]bool
	for i := 0; i < n/2; i++ {
		f[i+1][i] = true
		f[i][i+1] = true
	}
	if n%2 != 0 {
		f[0][0] = true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if f[i][j] {
				fmt.Fprintf(out, "%c", '#')
			} else {
				fmt.Fprintf(out, "%c", '.')
			}
		}
		fmt.Fprintln(out)
	}
}
