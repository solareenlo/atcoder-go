package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	flag := false
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a%(x+y) >= x {
			flag = true
		}
	}

	if !flag {
		fmt.Println("Second")
	} else if x <= y {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
