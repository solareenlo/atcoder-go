package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	p := getLine()
	a := len(p)
	if p[a-1] == 't' && p[a-2] == 'o' && p[a-3] == 'n' {
		fmt.Println(p)
		return
	}
	for i := 0; i < a; i++ {
		if p[i] == 'n' && p[i+1] == 'o' && p[i+2] == 't' && p[i+3] == ' ' && p[i+4] == 'n' && p[i+5] == 'o' && p[i+6] == 't' && p[i+7] == ' ' {
			i = i + 7
			continue
		}
		fmt.Printf("%c", p[i])
	}
	fmt.Println()
}

var in = bufio.NewScanner(os.Stdin)

func getLine() string {
	in.Scan()
	return in.Text()
}
