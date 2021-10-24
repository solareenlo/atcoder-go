package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	for i := 0; i < len(s); i++ {
		if i%2 != 0 {
			if s[i] == 'L' || s[i] == 'U' || s[i] == 'D' {
				cnt++
			}
		} else {
			if s[i] == 'R' || s[i] == 'U' || s[i] == 'D' {
				cnt++
			}
		}
	}

	if cnt == len(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
