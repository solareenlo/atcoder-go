package main

import "fmt"

func main() {
	var n int
	var t string
	fmt.Scan(&n, &t)

	s := "110"
	res := 0
	for i := 0; i < 3; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if t[j] != s[(j+i)%3] {
				ok = false
			}
		}
		if ok {
			res += (int(3e10)-n-i)/3 + 1
		}
	}
	fmt.Println(res)
}
