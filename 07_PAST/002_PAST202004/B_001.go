package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	t := make([]int, 3)
	ans := "abc"
	for i := range s {
		t[int(s[i]-'a')]++
	}

	for i := 0; i < 3; i++ {
		if t[i] > t[(i+1)%3] && t[i] > t[(i+2)%3] {
			fmt.Print(string(ans[i]))
		}
	}
}
