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

	if h%2 != 0 && w%2 != 0 {
		fmt.Println("First")
		fmt.Printf("%d %d %d\n", h/2+1, w/2+1, 1)
	} else {
		fmt.Println("Second")
	}

	var a, b, c int
	for {
		fmt.Fscan(in, &a, &b, &c)
		if a == -1 {
			break
		}
		tmp := c ^ 1
		if h%2 != 0 && w%2 != 0 {
			tmp = c
		}
		fmt.Printf("%d %d %d\n", h-a+1, w-b+1, tmp)
	}
}
