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

	ans := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans += (-n + 2*i + 1) * a
	}
	fmt.Println(ans)
}
