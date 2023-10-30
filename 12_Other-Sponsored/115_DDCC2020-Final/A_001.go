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
	s := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		t := 0
		for j := 2; j <= a; j++ {
			for a%j == 0 {
				a /= j
				t++
			}
		}
		s ^= t
	}
	if s != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
