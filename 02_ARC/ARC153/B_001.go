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
	s := make([]string, h)
	for t := range s {
		fmt.Fscan(in, &s[t])
	}
	var q int
	fmt.Fscan(in, &q)
	sgn := 1
	r := 0
	c := 0
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		r = (a*sgn + r + h) % h
		c = (b*sgn + c + w) % w
		sgn *= -1
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(string(s[(i*sgn+r+h)%h][(j*sgn+c+w)%w]))
		}
		fmt.Println()
	}
}
