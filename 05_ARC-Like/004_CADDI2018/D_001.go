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

	v := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		v |= a & 1
	}

	if v != 0 {
		fmt.Println("first")
	} else {
		fmt.Println("second")
	}
}
