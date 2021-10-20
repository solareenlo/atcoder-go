package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := map[byte]struct{}{}
	for i := 0; i < 4; i++ {
		m[s[i]] = struct{}{}
	}

	ok := true
	if len(m) != 2 {
		ok = false
	} else if s[0] == s[1] && s[1] == s[2] {
		ok = false
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
