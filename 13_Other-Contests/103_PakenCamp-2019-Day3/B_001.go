package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	b := 0
	w := 0
	for i := 1; i <= n; i++ {
		var c string
		fmt.Fscan(in, &c)
		if c[0] == 'b' {
			b++
		} else {
			w++
		}
	}
	if b > w {
		fmt.Println("black")
	} else {
		fmt.Println("white")
	}
}
