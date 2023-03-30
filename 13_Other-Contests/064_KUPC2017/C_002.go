package main

import "fmt"

func main() {
	var n byte
	var s = make([]byte, 1)
	fmt.Scan(&n, &s)

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
