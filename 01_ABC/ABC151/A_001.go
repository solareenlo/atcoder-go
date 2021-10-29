package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	alpha := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(string(alpha[c[0]-'a'+1]))
}
