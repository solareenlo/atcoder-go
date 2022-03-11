package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	d := make([]int, 5)
	for i := 1; i <= 6; i++ {
		var x int
		fmt.Fscan(in, &x)
		d[x]++
	}

	for i := 1; i <= 4; i++ {
		if d[i] > 2 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
