package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	s = getLine()

	var ans string
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			ans += "< "
		}
		if s[i] == 'R' {
			ans += "> "
		}
		if s[i] == 'A' {
			ans += "A "
		}
	}

	fmt.Println(ans[:len(ans)-1])
}

var in = bufio.NewScanner(os.Stdin)

func getLine() string {
	in.Scan()
	return in.Text()
}
