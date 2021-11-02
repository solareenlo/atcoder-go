package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	s1 := s[:len(s)/2]
	s2 := s[len(s)/2+1:]
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
