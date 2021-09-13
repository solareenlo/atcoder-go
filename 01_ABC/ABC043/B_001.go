package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != 'B' {
			res += string(s[i])
		} else if len(res) != 0 {
			res = res[:len(res)-1]
		}
	}
	fmt.Println(res)
}
