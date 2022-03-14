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

	m := map[string]int{}
	sum := 0
	for i := 0; i < n; i++ {
		var s string
		var t int
		fmt.Fscan(in, &s, &t)
		sum += t
		m[s] = sum
	}

	var x string
	fmt.Fscan(in, &x)
	fmt.Println(sum - m[x])
}
