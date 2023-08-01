package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	var x [110]string
	sum := 0
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &x[i])
		for j := 0; j < b; j++ {
			if x[i][j] != '.' {
				sum += int(x[i][j]) - 48
			}
		}
	}
	fmt.Println(sum)
}
