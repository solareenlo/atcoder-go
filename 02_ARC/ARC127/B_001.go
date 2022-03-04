package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l int
	fmt.Scan(&n, &l)

	s := [17]byte{}
	for k := 0; k < 3; k++ {
		for i := 0; i < n; i++ {
			s[0] = byte('0' + (2+k)%3)
			for j, x := l-1, i; j > 0; j-- {
				s[j] = byte('0' + (x+k)%3)
				x /= 3
			}
			for i := range s {
				if s[i] != 0 {
					fmt.Fprint(out, string(s[i]))
				}
			}
			fmt.Fprintln(out)
		}
	}
}
