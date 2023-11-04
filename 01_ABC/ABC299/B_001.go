package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var as [2000200]int

	var a, b int
	fmt.Fscan(in, &a, &b)
	for i := 1; i <= a; i++ {
		fmt.Fscan(in, &as[i])
	}
	w1, w2, m1, m2 := 0, 0, 0, 0
	for i := 1; i <= a; i++ {
		var c int
		fmt.Fscan(in, &c)
		if as[i] == b && c > m1 {
			m1 = c
			w1 = i
		}
		if as[i] == as[1] && c > m2 {
			m2 = c
			w2 = i
		}
	}
	if w1 != 0 {
		fmt.Println(w1)
	} else {
		fmt.Println(w2)
	}
}
