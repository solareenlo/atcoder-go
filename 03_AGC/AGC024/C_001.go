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

	res := -1
	now := -1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a > now+1 {
			fmt.Println(-1)
			return
		}
		if a > now {
			res++
		} else {
			res += a
		}
		now = a
	}

	fmt.Println(res)
}
