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

	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(in, &t)
		s[i] = strings.Split(t, "")
	}

	for up := -1; up <= 1; up++ {
		for lf := -1; lf <= 1; lf++ {
			if up == 0 && lf == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if s[i][j] == "Q" {
						x := i + up
						y := j + lf
						for 0 <= x && x < n && 0 <= y && y < n {
							if s[x][y] == "." {
								x += up
								y += lf
							} else {
								if s[x][y] == "X" {
									x -= up
									y -= lf
									s[x][y] = "#"
								}
								break
							}
						}
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, strings.Join(s[i], ""))
	}
}
