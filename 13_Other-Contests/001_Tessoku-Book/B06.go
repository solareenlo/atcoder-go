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
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(in, &v)
		s[i+1] = s[i] + v
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		A := s[r] - s[l]
		B := r - l - A
		if A == B {
			fmt.Println("draw")
		} else if A > B {
			fmt.Println("win")
		} else {
			fmt.Println("lose")
		}
	}
}
