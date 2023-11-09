package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [1000][1000]int

	var n int
	fmt.Fscan(in, &n)
	o := 0
	for n > 0 {
		n--
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		for x := a; x < b; x++ {
			for y := c; y < d; y++ {
				o += 1 - f[x][y]
				f[x][y] = 1
			}
		}
	}
	fmt.Println(o)
}
