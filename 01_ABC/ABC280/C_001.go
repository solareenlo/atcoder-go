package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, t string
	fmt.Fscan(in, &s, &t)

	s += "."
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			fmt.Println(i + 1)
			return
		}
	}
}
