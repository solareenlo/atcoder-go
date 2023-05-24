package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var H, W, N int
	fmt.Fscan(in, &H, &W, &N)
	m := make(map[pair]int)
	for N > 0 {
		N--
		var x, y int
		fmt.Fscan(in, &x, &y)
		m[pair{x, y}]++
		if x > 1 {
			m[pair{x - 1, y}]++
		}
		if y > 1 {
			m[pair{x, y - 1}]++
		}
		if x > 1 && y > 1 {
			m[pair{x - 1, y - 1}]++
		}
	}
	ans := 0
	for _, p := range m {
		ans += p & 1
	}
	fmt.Println(ans)
}
