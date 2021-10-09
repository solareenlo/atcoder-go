package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	for i := 0; i < len(s); i++ {
		s1 := s[:i]
		s2 := s[i:]
		if s2+s1 == t {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
