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

	c := make(map[int]int)
	c[0] = 1
	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		sum += x
		ans += c[sum]
		c[sum]++
	}
	fmt.Println(ans)
}
