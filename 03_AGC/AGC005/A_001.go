package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x string
	fmt.Fscan(in, &x)

	s, t := 0, 0
	for i := range x {
		if x[i] == 'S' {
			s++
		} else {
			if s != 0 {
				s--
			} else {
				t++
			}
		}
	}
	fmt.Println(s + t)
}
