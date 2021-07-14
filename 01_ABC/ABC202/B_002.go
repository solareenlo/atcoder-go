package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	s = reverse2(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			fmt.Print(9)
		} else if s[i] == '9' {
			fmt.Print(6)
		} else {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}

func reverse2(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return res
}
