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
	for i := 1; i <= n; i++ {
		var a, b, c string
		fmt.Fscan(in, &a, &b, &c)
		l := len(c)
		if a == "BEGINNING" {
			fmt.Printf("%c", c[0])
		} else if a == "MIDDLE" {
			fmt.Printf("%c", c[l>>1])
		} else {
			fmt.Printf("%c", c[l-1])
		}
	}
	fmt.Println()
}
