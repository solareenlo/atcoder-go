package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	fmt.Fscan(in, &n, &a)

	for i := 0; i < 4*n-2; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		a ^= tmp
	}
	fmt.Println(a)
}
