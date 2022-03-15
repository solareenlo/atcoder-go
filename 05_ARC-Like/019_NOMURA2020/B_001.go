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

	for i := range s {
		if s[i] != '?' {
			fmt.Print(string(s[i]))
		} else {
			fmt.Print("D")
		}
	}
}
