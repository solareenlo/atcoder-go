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

	c := make([]int, 2400025)
	for i := 0; i < N; i++ {
		var d, s, t int
		fmt.Fscan(in, &d, &s, &t)
		d *= 24
		s += d
		t += d
		c[s]++
		c[t]--
	}

	nc := 0
	for i := 0; i < 2400025; i++ {
		nc += c[i]
		if nc > 1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
