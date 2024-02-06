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
	sl := make(map[int]int)
	for i := 0; i < n; i++ {
		var s, c int
		fmt.Fscan(in, &s, &c)
		sl[s/(s&-s)] += c * (s & -s)
	}
	a := 0
	for _, v := range sl {
		for i := 0; i < 60; i++ {
			if (v>>i)&1 != 0 {
				a++
			}
		}
	}
	fmt.Println(a)
}
