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
	s := make([]int, 0)
	r := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		for len(s) > 0 && s[len(s)-1] > a {
			s = s[:len(s)-1]
		}
		if len(s) == 0 || s[len(s)-1] < a {
			s = append(s, a)
			r++
		}
	}
	fmt.Println(r)
}
