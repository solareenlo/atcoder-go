package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	cnt := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a == x {
			cnt++
		}
	}
	fmt.Println(cnt)
}
