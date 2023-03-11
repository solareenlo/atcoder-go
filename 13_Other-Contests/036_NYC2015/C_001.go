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
	if s[0] != t[0] {
		fmt.Println("No")
		return
	}

	var i int
	match := 0
	for i = 0; i < len(t); i++ {
		if match < len(s) && t[i] == s[match] {
			match++
		}
	}
	if match < len(s) {
		fmt.Println("No")
		return
	}

	for i = 0; i < len(t); i++ {
		if t[i] != t[0] {
			break
		}
	}

	tlen := i
	for i = 0; i < len(s); i++ {
		if s[i] != s[0] {
			break
		}
	}

	slen := i
	if slen < tlen {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
