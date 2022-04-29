package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	h, w int
	s    = make([][]string, 0)
)

func f(i, j int, c string) {
	if i < 0 || h <= i || j < 0 || w <= j {
		return
	}
	if s[i][j] == "#" {
		return
	}
	if s[i][j] != "." && s[i][j] != c {
		return
	}
	s[i][j] = "o"
	f(i+1, j, "^")
	f(i, j+1, "<")
	f(i-1, j, "v")
	f(i, j-1, ">")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var r, c int
	fmt.Fscan(in, &h, &w, &r, &c)
	r--
	c--
	s = make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(in, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	s[r][c] = "."
	f(r, c, ".")
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != "#" && s[i][j] != "o" {
				s[i][j] = "x"
			}
		}
		fmt.Fprintln(out, strings.Join(s[i], ""))
	}
}
