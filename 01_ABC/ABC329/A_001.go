package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s)-1; i++ {
		fmt.Printf("%s ", string(s[i]))
	}
	fmt.Println(string(s[len(s)-1]))
}
