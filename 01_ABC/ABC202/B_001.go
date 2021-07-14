package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	s = reverse(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			fmt.Print("9")
		} else if s[i] == '9' {
			fmt.Print("6")
		} else {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
