package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var d [1000]int

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x == 1 && d[y] != -1 {
			d[y]++
			if d[y] == 2 {
				d[y] = -1
			}
		} else if x == 2 && d[y] != -1 {
			d[y] = -1
		} else {
			if d[y] == -1 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
