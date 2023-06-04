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
	a := 0
	b := len(s)
	for a < b && s[a] == '_' {
		a++
	}
	for b > 0 && s[b-1] == '_' {
		b--
	}
	f := 0
	for i := a; i < b; i++ {
		if s[i] == '_' {
			f++
		}
	}
	if s == "_" {
		fmt.Print(s)
	} else if 'a' <= s[a] && s[a] <= 'z' && f == 0 {
		for i := 0; i < len(s); i++ {
			if 'A' <= s[i] && s[i] <= 'Z' {
				fmt.Print("_", string(s[i]+32))
			} else {
				fmt.Print(string(s[i]))
			}
		}
	} else {
		ng := false
		for i := a; i < b; i++ {
			if (s[i] == '_' && !('a' <= s[i+1] && s[i+1] <= 'z')) || ('A' <= s[i] && s[i] <= 'Z') {
				ng = true
			}
		}
		if 'a' <= s[a] && s[a] <= 'z' && !ng {
			for i := 0; i < len(s); i++ {
				if a <= i && i < b && s[i] == '_' {
					fmt.Print(string(s[i+1] - 32))
					i++
				} else {
					fmt.Print(string(s[i]))
				}
			}
		} else {
			fmt.Print(s)
		}
	}
	fmt.Println()
}
