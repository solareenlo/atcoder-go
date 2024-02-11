package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	chk := s[0] == '1'
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			chk = false
		}
	}

	if chk {
		fmt.Println(len(s) - 1)
	} else {
		fmt.Println(len(s))
	}
}
