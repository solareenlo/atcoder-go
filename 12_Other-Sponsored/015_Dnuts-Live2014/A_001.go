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

	if n%2 == 1 {
		fmt.Println("error")
		return
	}

	n /= 2
	ans := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		ans += abs(a - b)
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
