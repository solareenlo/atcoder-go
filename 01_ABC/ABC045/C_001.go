package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	res := 0
	for bit := 0; bit < 1<<(n-1); bit++ {
		tmp := ""
		for i := 0; i < n; i++ {
			tmp += string(s[i])
			if bit&(1<<i) == 0 {
				t, _ := strconv.Atoi(tmp)
				res += t
				tmp = ""
			}
		}
	}
	fmt.Println(res)
}
