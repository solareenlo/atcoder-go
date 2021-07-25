package main

import "fmt"

func main() {
	s := make(map[string]int)
	for i := 0; i < 4; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s[tmp]++
	}
	if len(s) == 4 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
