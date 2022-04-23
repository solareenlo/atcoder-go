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

	var n int
	fmt.Fscan(in, &n)
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var S string
		fmt.Fscan(in, &S)
		s[i] = strings.Split(S, "")
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 1 - i; j < n+i; j++ {
			for k := -1; k <= 1; k++ {
				if s[i+1][j+k] == "X" {
					s[i][j] = "X"
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, strings.Join(s[i], ""))
	}
}
