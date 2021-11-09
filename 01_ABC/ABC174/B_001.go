package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var d float64
	fmt.Fscan(in, &n, &d)

	cnt := 0
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		if x*x+y*y <= d*d {
			cnt++
		}
	}

	fmt.Println(cnt)
}
