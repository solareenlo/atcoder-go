package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var a int
	var s = make([]byte, 1)
	fmt.Scan(&a, &s)

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(a))
	n := b[0]

	for i := 0; i < len(s)*3; i++ {
		for j := len(s) - 2; j >= 0; j-- {
			if s[j+1] != 'a'-1 && s[j]+n <= 'z' {
				s[j] += n
				s[j+1]--
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
