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

	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			j--
			i++
		} else if s[j] != 'a' {
			fmt.Println("No")
			return
		} else {
			j--
		}
	}
	fmt.Println("Yes")
}
