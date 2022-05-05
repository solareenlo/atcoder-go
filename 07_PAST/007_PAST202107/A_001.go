package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ch := 0
	for i := 0; i < 14; i++ {
		if i%2 == 0 {
			ch += int(s[i]-'0') * 3
		} else {
			ch += int(s[i] - '0')
		}
	}

	if ch%10 == int(s[14]-'0') {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
