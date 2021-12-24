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

	ma, mb := 0, 1<<60
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if b < mb {
			mb = b
			ma = a
		}
	}

	fmt.Println(ma + mb)
}
