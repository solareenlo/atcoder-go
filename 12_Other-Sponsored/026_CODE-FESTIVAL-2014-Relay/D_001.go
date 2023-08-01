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
	b := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	x, y := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[i][j] == 'X' {
				x += i
			} else if b[i][j] == 'Y' {
				y += n - 1 - i
			}
		}
	}
	if x < y {
		fmt.Println("Y")
	} else if x > y {
		fmt.Println("X")
	} else {
		fmt.Println("Impossible")
	}
}
