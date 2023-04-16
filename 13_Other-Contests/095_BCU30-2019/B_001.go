package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	ans := 1
	before := -1
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a < before {
			ans++
		}
		before = a
	}

	if before == 0 {
		fmt.Println(ans - 1)
	} else {
		fmt.Println(ans)
	}
}
