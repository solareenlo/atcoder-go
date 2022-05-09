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

	cnt := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a = b - a
		a %= 100
		if a >= 50 {
			cnt++
		}
		a %= 10
		if a >= 5 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
