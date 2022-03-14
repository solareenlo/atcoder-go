package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &m)

	sum := 0
	ans := 0
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		sum += a * b
		ans += b
	}

	sum--
	ans--
	fmt.Println(sum/9 + ans)
}
