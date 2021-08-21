package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)
	for i := 0; i < len(w); i++ {
		if w[i] == 'a' || w[i] == 'i' || w[i] == 'u' || w[i] == 'e' || w[i] == 'o' {
			continue
		}
		fmt.Print(string(w[i]))
	}
	fmt.Println("")
}
