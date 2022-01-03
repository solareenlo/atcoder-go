package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	s := make([]string, l+1)
	for i := 0; i < l+1; i++ {
		s[i] = getLine()
	}

	var p int
	for i := 0; i < len(s[l]); i++ {
		if s[l][i] == 'o' {
			p = i
		}
	}

	for i := l; i >= 0; i-- {
		if p-1 >= 0 && s[i][p-1] == '-' {
			p -= 2
		} else if p+1 < len(s[i]) && s[i][p+1] == '-' {
			p += 2
		}
	}

	fmt.Println(p/2 + 1)
}

var in = bufio.NewScanner(os.Stdin)

func getLine() string {
	in.Scan()
	return in.Text()
}
