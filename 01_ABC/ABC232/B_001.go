package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	m := map[int]bool{}
	for i := range s {
		m[int(t[i]+26-s[i])%26] = true
	}

	if len(m) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
